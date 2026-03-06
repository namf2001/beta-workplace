package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
// @Failure      400  {object} response.Error
// @Failure      409  {object} response.Error
// @Failure      500  {object} response.Error
// @Security     BearerAuth
// @Router       /users [post]
func (h Handler) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			response.HandleError(c, err)
			return
		}

		if err := validator.Validate(req); err != nil {
			response.HandleError(c, webErrValidationFailed)
			return
		}

		input := ctrlUsers.CreateUserInput{
			Email: req.Email,
			Name:  req.Name,
		}

		user, err := h.userCtrl.CreateUser(c.Request.Context(), input)
		if err != nil {
			response.HandleError(c, convertError(err))
			return
		}

		c.JSON(http.StatusCreated, CreateUserResponse{User: user})
	}
}
