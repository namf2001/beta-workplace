package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/namf2001/beta-workplace/constants"
	"github.com/namf2001/beta-workplace/internal/handler/response"
)

// GetProfileResponse is DTO for GetProfile response
type GetProfileResponse struct {
	ID            int64     `json:"id"`
	Email         string    `json:"email"`
	FullName      string    `json:"full_name"`
	ProfileImage  string    `json:"profile_image"`
	EmailVerified time.Time `json:"email_verified"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// GetProfile handles retrieving the current user's profile
// @Summary      Get user profile
// @Description  Get the profile of the currently authenticated user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200  {object} response.Response
// @Failure      401  {object} response.Response
// @Failure      404  {object} response.Response
// @Security     BearerAuth
// @Router       /auth/user/profile [get]
func (h Handler) GetProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, response.NewResponse(
				constants.InvalidToken.Code,
				constants.InvalidToken.Message,
				nil,
			))
			return
		}

		user, err := h.ctrl.GetUserProfile(c.Request.Context(), userID.(int64))
		if err != nil {
			c.JSON(http.StatusNotFound, response.NewResponse(
				constants.UserNotFound.Code,
				err.Error(),
				nil,
			))
			return
		}

		c.JSON(http.StatusOK, response.NewResponse(
			constants.GetUserInfoSuccess.Code,
			constants.GetUserInfoSuccess.Message,
			user,
		))
	}
}
