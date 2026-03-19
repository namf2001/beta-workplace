package auth

import "context"

func (i impl) Logout(ctx context.Context, userID int64) error {
	// Right now JWT is stateless, so we just return success.
	// If we tracked session inside sessions table, we'd delete it here.
	return nil
}
