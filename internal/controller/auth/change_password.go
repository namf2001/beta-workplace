package auth

import (
	"context"

	"github.com/namf2001/beta-workplace/internal/pkg/utils"
)

// ChangePassword changes user password
func (i impl) ChangePassword(ctx context.Context, userID int64, oldPassword, newPassword string) error {
	// Get current user
	user, err := i.repo.User().GetByID(ctx, userID)
	if err != nil {
		return err
	}

	// Verify old password
	if err := utils.VerifyPassword(user.Password, oldPassword); err != nil {
		return err
	}

	// Hash new password
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return err
	}

	// Update password in database
	if err := i.repo.User().UpdatePassword(ctx, userID, hashedPassword); err != nil {
		return err
	}

	return nil
}
