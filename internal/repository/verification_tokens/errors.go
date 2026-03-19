package verification_tokens

import "errors"

var (
	// ErrTokenNotFoundOrExpired is returned when a token is invalid or expired
	ErrTokenNotFoundOrExpired = errors.New("verification token not found or already expired")
)
