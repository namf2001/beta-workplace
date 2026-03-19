package users

import (
	"context"
	"database/sql"

	"github.com/namf2001/beta-workplace/internal/model"
	"github.com/namf2001/beta-workplace/internal/model/common"
	pkgerrors "github.com/pkg/errors"
)

// GetByEmail implements Repository.
func (i impl) GetByEmail(ctx context.Context, email string) (model.User, error) {
	query := `
		SELECT id, email, name, password, image, email_verified, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	var user model.User
	var ev sql.NullTime
	err := i.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.Password,
		&user.Image,
		&ev,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if ev.Valid {
		user.EmailVerified = &ev.Time
	}

	if err == sql.ErrNoRows {
		return model.User{}, pkgerrors.WithStack(common.ErrUserNotFound)
	}

	if err != nil {
		return model.User{}, pkgerrors.WithStack(err)
	}

	return user, nil
}
