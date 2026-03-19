package verification_tokens

import (
	"context"

	"github.com/namf2001/beta-workplace/internal/model"
	pkgerrors "github.com/pkg/errors"
)

// Create saves a new verification token
func (i *impl) Create(ctx context.Context, token model.VerificationToken) error {
	const query = `
		INSERT INTO verification_token (
			identifier,
			expires,
			token
		) VALUES (
			$1, $2, $3
		)
	`

	_, err := i.db.ExecContext(ctx, query,
		token.Identifier,
		token.Expires,
		token.Token,
	)

	return pkgerrors.WithStack(err)
}
