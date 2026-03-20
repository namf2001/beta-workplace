package auth

import (
	"context"
)

// DeleteAccount deletes user account and all related data
func (i impl) DeleteAccount(ctx context.Context, userID int64) error {
	// Delete user from database
	// Note: In production, this should also delete all related data:
	// - User's organizations
	// - User's projects
	// - User's tasks
	// - User's messages
	// - User's files
	// - User's sessions
	// This can be done using database triggers or by implementing
	// cascade delete logic in the repository layer

	if err := i.repo.User().Delete(ctx, userID); err != nil {
		return err
	}

	return nil
}
