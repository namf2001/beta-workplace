package auth

import (
	"context"

	"github.com/namf2001/beta-workplace/internal/repository"
)

type Controller interface {
	// Login handles manual login
	Login(ctx context.Context, input ValidationInput) (string, error)

	// RegisterStep1SendOTP starts the registration process by sending an OTP
	RegisterStep1SendOTP(ctx context.Context, email string) error

	// RegisterStep2VerifyOTP verifies the OTP for registration
	RegisterStep2VerifyOTP(ctx context.Context, email, otp string) error

	// RegisterStep3Complete completes the registration and returns a JWT token
	RegisterStep3Complete(ctx context.Context, input RegisterInput) (string, error)

	// ForgotPasswordStep1SendOTP sends an OTP for password reset
	ForgotPasswordStep1SendOTP(ctx context.Context, email string) error

	// ForgotPasswordStep2Reset verifies OTP and resets password
	ForgotPasswordStep2Reset(ctx context.Context, email, otp, newPassword string) error

	// Logout handles user logout
	Logout(ctx context.Context, userID int64) error

	// OAuthLogin handles oauth login/registration
	OAuthLogin(ctx context.Context, input OAuthInput) (string, error)
}

type impl struct {
	repo repository.Registry
}

func New(repo repository.Registry) Controller {
	return impl{
		repo: repo,
	}
}
