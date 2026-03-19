package users

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/namf2001/beta-workplace/constants"
	ctrlUsers "github.com/namf2001/beta-workplace/internal/controller/users"
	"github.com/namf2001/beta-workplace/internal/handler/response"
	"github.com/namf2001/beta-workplace/internal/model"
	"github.com/namf2001/beta-workplace/internal/pkg/validator"
)

// CreateUserRequest represents the request for creating a user
type CreateUserRequest struct {
	Email string `json:"email" validate:"required,email"`
	Name  string `json:"name" validate:"required,min=2,max=100"`
}

// CreateUserResponse represents the response for creating a user
type CreateUserResponse struct {
	User model.User `json:"user"`
}

// CreateUser handles the creation of a new user
// @Summary      Create user
// @Description  Create a new user account
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        input body users.CreateUserRequest true "User info"
// @Success      201  {object} users.CreateUserResponse
// @Failure      400  {object} response.Response
// @Failure      409  {object} response.Response
// @Failure      500  {object} response.Response
// @Security     BearerAuth
// @Router       /users [post]
func (h Handler) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateUserRequest
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

		input := ctrlUsers.CreateUserInput{
			Email: req.Email,
			Name:  req.Name,
		}

		user, err := h.userCtrl.CreateUser(c.Request.Context(), input)
		if err != nil {
			code := constants.InternalServerError.Code
			if errors.Is(err, ctrlUsers.ErrUserExited) {
				code = constants.EmailExists.Code
			}
			c.JSON(http.StatusInternalServerError, response.NewResponse(
				code,
				err.Error(),
				nil,
			))
			return
		}

		c.JSON(http.StatusCreated, response.NewResponse(
			constants.RegisterUserSuccess.Code,
			constants.RegisterUserSuccess.Message,
			CreateUserResponse{User: user},
		))
	}
}
