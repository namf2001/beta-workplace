package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/namf2001/beta-workplace/constants"
	"github.com/namf2001/beta-workplace/internal/handler/response"
	"github.com/namf2001/beta-workplace/internal/pkg/jwt"
)

// Logout handles user logout
// @Summary      Logout user
// @Description  Logout the user by clearing their session/token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200  {object} response.Response
// @Failure      401  {object} response.Response
// @Failure      500  {object} response.Response
// @Security     BearerAuth
// @Router       /auth/logout [post]
func (h *Handler) Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		claimsVal, exists := c.Get("user_claims")
		if !exists {
			c.JSON(http.StatusInternalServerError, response.NewResponse(
				constants.LogoutFail.Code,
				constants.LogoutFail.Message,
				nil,
			))
			return
		}

		claims, ok := claimsVal.(*jwt.Claims)
		if !ok {
			c.JSON(http.StatusInternalServerError, response.NewResponse(
				constants.LogoutFail.Code,
				constants.LogoutFail.Message,
				nil,
			))
			return
		}

		err := h.ctrl.Logout(c.Request.Context(), claims.UserID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.NewResponse(
				constants.InternalServerError.Code,
				err.Error(),
				nil,
			))
			return
		}

		c.JSON(http.StatusOK, response.NewResponse(
			constants.LogoutSuccess.Code,
			constants.LogoutSuccess.Message,
			nil,
		))
	}
}
