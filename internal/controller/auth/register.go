package auth

import (
	"context"
	"time"

	"github.com/namf2001/beta-workplace/internal/model"
	"github.com/namf2001/beta-workplace/internal/pkg/jwt"
	"github.com/namf2001/beta-workplace/internal/pkg/mail"
	"github.com/namf2001/beta-workplace/internal/pkg/utils"
	"github.com/namf2001/beta-workplace/internal/repository"
)

type RegisterInput struct {
	Name     string
	Email    string
	OTP      string
	Password string
}

func (i impl) RegisterStep1SendOTP(ctx context.Context, email string) error {
	// 1. Check if email already registered
	existingUser, _ := i.repo.User().GetByEmail(ctx, email)
	if existingUser.ID != 0 {
		return ErrEmailAlreadyExists
	}

	// 2. Clear old OTPs
	err := i.repo.VerificationToken().DeleteAllForIdentifier(ctx, email)
	if err != nil {
		return err
	}

	// 3. Generate new OTP
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

	// 5. Send Email
	return mail.SendMailRegistration("User", email, otp)
}

func (i impl) RegisterStep2VerifyOTP(ctx context.Context, email, otp string) error {
	_, err := i.repo.VerificationToken().GetValidToken(ctx, email, otp)
	if err != nil {
		return ErrWrongOTP
	}
	return nil
}

func (i impl) RegisterStep3Complete(ctx context.Context, input RegisterInput) (string, error) {
	// 1. Verify OTP again
	_, err := i.repo.VerificationToken().GetValidToken(ctx, input.Email, input.OTP)
	if err != nil {
		return "", ErrWrongOTP
	}

	// 2. Hash password
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return "", err
	}

	// 3. Create user + account
	user := model.User{
		Name:          input.Name,
		Email:         input.Email,
		Password:      hashedPassword,
		EmailVerified: time.Now(),
	}

	var createdUser model.User
	err = i.repo.DoInTx(ctx, func(ctx context.Context, txRepo repository.Registry) error {
		var txErr error
		createdUser, txErr = txRepo.User().Create(ctx, user)
		if txErr != nil {
			return txErr
		}

		_, txErr = txRepo.Account().Create(ctx, model.Account{
			UserID:            createdUser.ID,
			Type:              "credentials",
			Provider:          model.ProviderCredentials,
			ProviderAccountID: input.Email,
		})
		return txErr
	}, nil)
	if err != nil {
		return "", err
	}

	// 4. Delete OTP after successful registration
	err = i.repo.VerificationToken().DeleteAllForIdentifier(ctx, input.Email)
	if err != nil {
		return "", err
	}

	// 5. Generate token
	token, err := jwt.GenerateToken(createdUser.ID, createdUser.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
