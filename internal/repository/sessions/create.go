package sessions

import (
	"context"

	"github.com/namf2001/beta-workplace/internal/model"
	pkgerrors "github.com/pkg/errors"
)

// Create implements Repository.
func (i impl) Create(ctx context.Context, session model.Session) (model.Session, error) {
	query := `
		INSERT INTO sessions (user_id, expires, session_token)
		VALUES ($1, $2, $3)
		RETURNING id, user_id, expires, session_token
	`

	var created model.Session
	err := i.db.QueryRowContext(ctx, query, session.UserID, session.Expires, session.SessionToken).Scan(
		&created.ID,
		&created.UserID,
		&created.Expires,
		&created.SessionToken,
	)

	if err != nil {
		return model.Session{}, pkgerrors.WithStack(err)
	}

	return created, nil
}
