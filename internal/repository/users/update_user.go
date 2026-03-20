package users

import (
	"context"
	"errors"

	"github.com/lib/pq"
	"github.com/namf2001/beta-workplace/internal/model"
	pkgerrors "github.com/pkg/errors"
)

// Update implements Repository.
func (i impl) Update(ctx context.Context, user model.User) error {
	query := `
		UPDATE users
		SET email = $1, name = $2, password = $3, image = $4, email_verified = $5, updated_at = NOW()
		WHERE id = $6
	`

	result, err := i.db.ExecContext(ctx, query, user.Email, user.Name, user.Password, user.Image, user.EmailVerified, user.ID)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			return pkgerrors.WithStack(ErrDuplicateEmail)
		}
		return pkgerrors.WithStack(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return pkgerrors.WithStack(err)
	}

	if rowsAffected == 0 {
		return pkgerrors.WithStack(ErrNotFound)
	}

	return nil
}

// UpdatePassword updates user password
func (i impl) UpdatePassword(ctx context.Context, id int64, hashedPassword string) error {
	query := `
		UPDATE users
		SET password = $1, updated_at = NOW()
		WHERE id = $2
	`

	result, err := i.db.ExecContext(ctx, query, hashedPassword, id)
	if err != nil {
		return pkgerrors.WithStack(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return pkgerrors.WithStack(err)
	}

	if rowsAffected == 0 {
		return pkgerrors.WithStack(ErrNotFound)
	}

	return nil
}
