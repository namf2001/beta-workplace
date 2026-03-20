package auth

import (
	"context"

	"github.com/namf2001/beta-workplace/internal/model"
)

// UpdateUserProfile updates user profile
func (i impl) UpdateUserProfile(ctx context.Context, userID int64, input interface{}) (interface{}, error) {
	// Get current user
	user, err := i.repo.User().GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// Update user fields based on input
	// Note: In production, you should properly handle RSA decryption for sensitive fields
	// as mentioned in the API documentation
	if updateReq, ok := input.(map[string]interface{}); ok {
		if fullName, exists := updateReq["full_name"]; exists && fullName != "" {
			user.Name = fullName.(string)
		}
		if email, exists := updateReq["email"]; exists && email != "" {
			user.Email = email.(string)
		}
		if profileImage, exists := updateReq["profile_image"]; exists && profileImage != "" {
			user.Image = profileImage.(string)
		}
	}

	// Update user in database
	updatedUser := model.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Image: user.Image,
	}

	if err := i.repo.User().Update(ctx, updatedUser); err != nil {
		return nil, err
	}

	// Return updated user
	return map[string]interface{}{
		"id":        updatedUser.ID,
		"full_name": updatedUser.Name,
		"email":     updatedUser.Email,
	}, nil
}
