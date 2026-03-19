package utils

import (
	"crypto/rand"
	"math/big"

	pkgerrors "github.com/pkg/errors"
)

// GenerateOTP creates a random numeric OTP of the specified length
func GenerateOTP(length int) (string, error) {
	const charset = "0123456789"
	b := make([]byte, length)
	for i := range b {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", pkgerrors.WithStack(err)
		}
		b[i] = charset[n.Int64()]
	}
	return string(b), nil
}
