package auth

import (
	"context"
)

// GetUserProfile retrieves user profile by user ID
func (i impl) GetUserProfile(ctx context.Context, userID int64) (interface{}, error) {
	user, err := i.repo.User().GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// Return user without password
	return map[string]interface{}{
		"id":              user.ID,
		"full_name":       user.Name,
		"email":           user.Email,
		"profile_image":   user.Image,
		"email_verified":  user.EmailVerified,
		"created_at":      user.CreatedAt,
		"updated_at":      user.UpdatedAt,
	}, nil
}
