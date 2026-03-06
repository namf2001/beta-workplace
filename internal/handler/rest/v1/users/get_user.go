package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/namf2001/beta-workplace/internal/handler/response"
	"github.com/namf2001/beta-workplace/internal/model"
)

// GetUserResponse represents the response for getting a user
type GetUserResponse struct {
	User model.User `json:"user"`
}

// GetUser handles the retrieval of a user by ID
// @Summary      Get user
// @Description  Get user details by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object} users.GetUserResponse
// @Failure      400  {object} response.Error
// @Failure      404  {object} response.Error
// @Failure      500  {object} response.Error
// @Security     BearerAuth
// @Router       /users/{id} [get]
func (h Handler) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			response.HandleError(c, webErrInvalidID)
			return
		}

		user, err := h.userCtrl.GetUser(c.Request.Context(), id)
		if err != nil {
			response.HandleError(c, convertError(err))
			return
		}

		c.JSON(http.StatusOK, GetUserResponse{User: user})
	}
}
