package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/namf2001/beta-workplace/internal/handler/response"
)

// DeleteUser handles the deletion of a user by ID
// @Summary      Delete user
// @Description  Delete a user account
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      204  {object} nil
// @Failure      400  {object} response.Error
// @Failure      500  {object} response.Error
// @Security     BearerAuth
// @Router       /users/{id} [delete]
func (h Handler) DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			response.HandleError(c, webErrInvalidID)
			return
		}

		if err := h.userCtrl.DeleteUser(c.Request.Context(), id); err != nil {
			response.HandleError(c, convertError(err))
			return
		}

		c.Status(http.StatusNoContent)
	}
}
