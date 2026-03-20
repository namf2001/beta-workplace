package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/namf2001/beta-workplace/constants"
	"github.com/namf2001/beta-workplace/internal/handler/response"
	"github.com/namf2001/beta-workplace/internal/pkg/validator"
)

// UpdateProfileRequest represents the request for updating user profile
type UpdateProfileRequest struct {
	FullName     string     `json:"full_name" validate:"omitempty,min=2,max=100"`
	Email        string     `json:"email" validate:"omitempty,email"`
	PhoneNumber  string     `json:"phone_number" validate:"omitempty"`
	Address      string     `json:"address" validate:"omitempty"`
	DateOfBirth  *time.Time `json:"date_of_birth" validate:"omitempty"`
	Gender       *int       `json:"gender" validate:"omitempty,min=0,max=2"`
	Colour       string     `json:"colour" validate:"omitempty"`
	ProfileImage string     `json:"profile_image" validate:"omitempty,url"`
}

// UpdateProfile handles updating the current user's profile
// @Summary      Update user profile
// @Description  Update the profile of the currently authenticated user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body      auth.UpdateProfileRequest  true  "Update info"
// @Success      200  {object} response.Response
// @Failure      400  {object} response.Response
// @Failure      401  {object} response.Response
// @Failure      500  {object} response.Response
// @Security     BearerAuth
// @Router       /auth/user/profile [put]
func (h Handler) UpdateProfile() gin.HandlerFunc {
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

		var req UpdateProfileRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, response.NewResponse(
				constants.BindJSONFail.Code,
				constants.BindJSONFail.Message,
				nil,
			))
			return
		}

		if err := validator.Validate(req); err != nil {
			c.JSON(http.StatusBadRequest, response.NewResponse(
				constants.InvalidRequestParams.Code,
				constants.InvalidRequestParams.Message,
				nil,
			))
			return
		}

		user, err := h.ctrl.UpdateUserProfile(c.Request.Context(), userID.(int64), req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.NewResponse(
				constants.InternalServerError.Code,
				err.Error(),
				nil,
			))
			return
		}

		c.JSON(http.StatusOK, response.NewResponse(
			constants.UpdateProfileSuccess.Code,
			constants.UpdateProfileSuccess.Message,
			user,
		))
	}
}
