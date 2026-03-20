package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/namf2001/beta-workplace/constants"
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
// @Failure      400  {object} response.Response
// @Failure      500  {object} response.Response
// @Security     BearerAuth
// @Router       /users/{id} [delete]
func (h Handler) DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.NewResponse(
				constants.InvalidRequestParams.Code,
				constants.InvalidRequestParams.Message,
				nil,
			))
			return
		}

		if err := h.userCtrl.DeleteUser(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusInternalServerError, response.NewResponse(
				constants.DeleteUserFail.Code,
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
