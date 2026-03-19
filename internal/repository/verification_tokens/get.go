package verification_tokens

import (
	"context"
	"database/sql"
	"time"

	"github.com/namf2001/beta-workplace/internal/model"
	pkgerrors "github.com/pkg/errors"
)

// GetValidToken retrieves a valid verification token by identifier and token string
func (i *impl) GetValidToken(ctx context.Context, identifier, tokenStr string) (model.VerificationToken, error) {
	const query = `
		SELECT
			identifier,
			expires,
			token
		FROM verification_token
		WHERE identifier = $1 AND token = $2 AND expires > $3
	`

	var token model.VerificationToken
	err := i.db.QueryRowContext(ctx, query, identifier, tokenStr, time.Now()).Scan(
		&token.Identifier,
		&token.Expires,
		&token.Token,
	)

	if err == sql.ErrNoRows {
		return model.VerificationToken{}, pkgerrors.WithStack(ErrTokenNotFoundOrExpired)
	}

	return token, pkgerrors.WithStack(err)
}
