package auth

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrWrongOTP           = errors.New("wrong or expired OTP")
)
