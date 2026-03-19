package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/namf2001/beta-workplace/constants"
	ctrlUsers "github.com/namf2001/beta-workplace/internal/controller/users"
	"github.com/namf2001/beta-workplace/internal/handler/response"
	"github.com/namf2001/beta-workplace/internal/pkg/validator"
)

// UpdateUserRequest represents the request for updating a user
type UpdateUserRequest struct {
	Email string `json:"email" validate:"omitempty,email"`
	Name  string `json:"name" validate:"omitempty,min=2,max=100"`
}

// UpdateUser handles the updating of a user by ID
// @Summary      Update user
// @Description  Update user details
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id     path      int                    true  "User ID"
// @Param        input  body      users.UpdateUserRequest  true  "Update info"
// @Success      204  {object} nil
// @Failure      400  {object} response.Response
// @Failure      500  {object} response.Response
// @Security     BearerAuth
// @Router       /users/{id} [put]
func (h Handler) UpdateUser() gin.HandlerFunc {
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

		var req UpdateUserRequest
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

		input := ctrlUsers.UpdateUserInput{
			Email: req.Email,
			Name:  req.Name,
		}

		if err := h.userCtrl.UpdateUser(c.Request.Context(), id, input); err != nil {
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
			nil,
		))
	}
}
