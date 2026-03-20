package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/namf2001/beta-workplace/constants"
	"github.com/namf2001/beta-workplace/internal/handler/response"
)

// DeleteAccount handles deleting the current user's account
// @Summary      Delete account
// @Description  Delete the account of the currently authenticated user and all related data
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200  {object} response.Response
// @Failure      401  {object} response.Response
// @Failure      500  {object} response.Response
// @Security     BearerAuth
// @Router       /auth/user/account [delete]
func (h Handler) DeleteAccount() gin.HandlerFunc {
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

		err := h.ctrl.DeleteAccount(c.Request.Context(), userID.(int64))
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.NewResponse(
				constants.InternalServerError.Code,
				err.Error(),
				nil,
			))
			return
		}

		c.JSON(http.StatusOK, response.NewResponse(
			constants.DeleteUserSuccess.Code,
			constants.DeleteUserSuccess.Message,
			nil,
		))
	}
}
