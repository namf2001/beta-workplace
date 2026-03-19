package auth

import (
	"context"
	"time"

	"github.com/namf2001/beta-workplace/internal/model"
	"github.com/namf2001/beta-workplace/internal/pkg/mail"
	"github.com/namf2001/beta-workplace/internal/pkg/utils"
)

func (i impl) ForgotPasswordStep1SendOTP(ctx context.Context, email string) error {
	// 1. Check if user exists
	user, err := i.repo.User().GetByEmail(ctx, email)
	if err != nil {
		// Do not leak if user exists or not
		return nil
	}

	// 2. Clear old OTPs
	err = i.repo.VerificationToken().DeleteAllForIdentifier(ctx, email)
	if err != nil {
		return err
	}

	// 3. Generate OTP
	otp, err := utils.GenerateOTP(6)
	if err != nil {
		return err
	}

	// 4. Save to DB
	err = i.repo.VerificationToken().Create(ctx, model.VerificationToken{
		Identifier: email,
		Token:      otp,
		Expires:    time.Now().Add(15 * time.Minute),
	})
	if err != nil {
		return err
	}

	// 5. Send email
	return mail.SendMailForgotPassword(user.Name, email, otp)
}

func (i impl) ForgotPasswordStep2Reset(ctx context.Context, email, otp, newPassword string) error {
	// 1. Get User
	user, err := i.repo.User().GetByEmail(ctx, email)
	if err != nil {
		return ErrWrongOTP // Keep error generic to avoid email enumeration
	}

	// 2. Verify OTP
	_, err = i.repo.VerificationToken().GetValidToken(ctx, email, otp)
	if err != nil {
		return ErrWrongOTP
	}

	// 3. Hash new password
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return err
	}

	// 4. Update User Profile
	user.Password = hashedPassword
	err = i.repo.User().Update(ctx, user)
	if err != nil {
		return err
	}

	// 5. Delete OTP
	return i.repo.VerificationToken().DeleteAllForIdentifier(ctx, email)
}
