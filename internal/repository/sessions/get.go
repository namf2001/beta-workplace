package sessions

import (
	"context"
	"database/sql"

	"github.com/namf2001/beta-workplace/internal/model"
	pkgerrors "github.com/pkg/errors"
)

// GetByToken implements Repository.
func (i impl) GetByToken(ctx context.Context, token string) (model.Session, error) {
	query := `
		SELECT id, user_id, expires, session_token
		FROM sessions
		WHERE session_token = $1
	`

	var session model.Session
	err := i.db.QueryRowContext(ctx, query, token).Scan(
		&session.ID,
		&session.UserID,
		&session.Expires,
		&session.SessionToken,
	)

	if err == sql.ErrNoRows {
		return model.Session{}, pkgerrors.WithStack(ErrNotFound)
	}

	if err != nil {
		return model.Session{}, pkgerrors.WithStack(err)
	}

	return session, nil
}
