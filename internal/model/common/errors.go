package common

import "errors"

var (
	ErrInvalidEmail = errors.New("invalid email format")
	ErrInvalidName  = errors.New("name must be between 2 and 100 characters")
	ErrUserNotFound = errors.New("user not found")
)
