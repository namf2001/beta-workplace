package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/namf2001/beta-workplace/constants"
	ctrlUsers "github.com/namf2001/beta-workplace/internal/controller/users"
	"github.com/namf2001/beta-workplace/internal/handler/response"
	"github.com/namf2001/beta-workplace/internal/model"
)

// ListUsersRequest represents the request for listing users
type ListUsersRequest struct {
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Email  string `json:"email"`
}

// ListUsersResponse represents the response for listing users
type ListUsersResponse struct {
	Users  []model.User `json:"users"`
	Total  int64        `json:"total"`
	Limit  int          `json:"limit"`
	Offset int          `json:"offset"`
}

// ListUsers handles the listing of users with optional filters
// @Summary      List users
// @Description  Get a list of users
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        limit  query     int     false  "Limit"
// @Param        offset query     int     false  "Offset"
// @Param        email  query     string  false  "Email filter"
// @Success      200  {object} users.ListUsersResponse
// @Failure      500  {object} response.Response
// @Security     BearerAuth
// @Router       /users [get]
func (h Handler) ListUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse query parameters
		limitStr := c.Query("limit")
		offsetStr := c.Query("offset")
		email := c.Query("email")

		limit := 10 // default
		if limitStr != "" {
			if l, err := strconv.Atoi(limitStr); err == nil {
				limit = l
			}
		}

		offset := 0
		if offsetStr != "" {
			if o, err := strconv.Atoi(offsetStr); err == nil {
				offset = o
			}
		}

		filters := ctrlUsers.ListFilters{
			Limit:  limit,
			Offset: offset,
			Email:  email,
		}

		result, totalUser, err := h.userCtrl.ListUsers(c.Request.Context(), filters)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.NewResponse(
				constants.InternalServerError.Code,
				err.Error(),
				nil,
			))
			return
		}

		c.JSON(http.StatusOK, response.NewResponse(
			constants.GetUserInfoSuccess.Code,
			constants.GetUserInfoSuccess.Message,
			ListUsersResponse{
				Users:  result,
				Total:  totalUser,
				Limit:  limit,
				Offset: offset,
			},
		))
	}
}
