package verification_tokens

import (
	"context"

	pkgerrors "github.com/pkg/errors"
)

// Delete removes a token after it's been used
func (i *impl) Delete(ctx context.Context, identifier, tokenStr string) error {
	const query = `
		DELETE FROM verification_token
		WHERE identifier = $1 AND token = $2
	`

	_, err := i.db.ExecContext(ctx, query, identifier, tokenStr)
	return pkgerrors.WithStack(err)
}

// DeleteAllForIdentifier removes all tokens for a specific email/identifier
func (i *impl) DeleteAllForIdentifier(ctx context.Context, identifier string) error {
	const query = `
		DELETE FROM verification_token
		WHERE identifier = $1
	`

	_, err := i.db.ExecContext(ctx, query, identifier)
	return pkgerrors.WithStack(err)
}
