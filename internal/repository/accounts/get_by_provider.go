package accounts

import (
	"context"
	"database/sql"

	"github.com/namf2001/beta-workplace/internal/model"
	pkgerrors "github.com/pkg/errors"
)

// GetByProvider implements Repository.
func (i impl) GetByProvider(ctx context.Context, provider model.Provider, providerAccountID string) (model.Account, error) {
	query := `
		SELECT id, user_id, type, provider, provider_account_id, refresh_token, access_token, expires_at, id_token, scope, session_state, token_type
		FROM accounts
		WHERE provider = $1 AND provider_account_id = $2
	`

	var account model.Account
	err := i.db.QueryRowContext(ctx, query, provider, providerAccountID).Scan(
		&account.ID,
		&account.UserID,
		&account.Type,
		&account.Provider,
		&account.ProviderAccountID,
		&account.RefreshToken,
		&account.AccessToken,
		&account.ExpiresAt,
		&account.IDToken,
		&account.Scope,
		&account.SessionState,
		&account.TokenType,
	)

	if err == sql.ErrNoRows {
		return model.Account{}, pkgerrors.WithStack(ErrNotFound)
	}

	if err != nil {
		return model.Account{}, pkgerrors.WithStack(err)
	}

	return account, nil
}
