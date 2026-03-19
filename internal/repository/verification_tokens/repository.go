package verification_tokens

import (
	"context"

	"github.com/namf2001/beta-workplace/internal/model"
)

// Repository is the interface for verification token database operations
type Repository interface {
	// Create saves a new verification token
	Create(ctx context.Context, token model.VerificationToken) error

	// GetValidToken retrieves a valid verification token by identifier and token string
	GetValidToken(ctx context.Context, identifier, tokenStr string) (model.VerificationToken, error)

	// Delete removes a token after it's been used
	Delete(ctx context.Context, identifier, tokenStr string) error

	// DeleteAllForIdentifier removes all tokens for a specific email/identifier
	DeleteAllForIdentifier(ctx context.Context, identifier string) error
}
