package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/namf2001/beta-workplace/constants"
	"github.com/namf2001/beta-workplace/internal/handler/response"
	"github.com/namf2001/beta-workplace/internal/pkg/validator"
)

// ChangePasswordRequest represents the request for changing password
type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required,min=6"`
	NewPassword string `json:"new_password" validate:"required,min=6"`
}

// ChangePassword handles changing the current user's password
// @Summary      Change password
// @Description  Change the password of the currently authenticated user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body      auth.ChangePasswordRequest  true  "Password info"
// @Success      200  {object} response.Response
// @Failure      400  {object} response.Response
// @Failure      401  {object} response.Response
// @Failure      500  {object} response.Response
// @Security     BearerAuth
// @Router       /auth/user/password [patch]
func (h Handler) ChangePassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get userID from context (set by auth middleware)
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, response.NewResponse(
				constants.InvalidToken.Code,
				constants.InvalidToken.Message,
				nil,
			))
			return
		}

		var req ChangePasswordRequest
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

		err := h.ctrl.ChangePassword(c.Request.Context(), userID.(int64), req.OldPassword, req.NewPassword)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.NewResponse(
				constants.PasswordIncorrect.Code,
				err.Error(),
				nil,
			))
			return
		}

		c.JSON(http.StatusOK, response.NewResponse(
			constants.ChangePasswordSuccess.Code,
			constants.ChangePasswordSuccess.Message,
			nil,
		))
	}
}
