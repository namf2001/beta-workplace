package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ctrlAuth "github.com/namf2001/beta-workplace/internal/controller/auth"
	"github.com/namf2001/beta-workplace/internal/handler/response"
	"github.com/namf2001/beta-workplace/internal/pkg/validator"
)

type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type RegisterResponse struct {
	Token string `json:"token"`
}

// Register handles manual registration
// @Summary      Register user
// @Description  Register a new user account
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input body auth.RegisterRequest true "Registration info"
// @Success      200  {object} auth.RegisterResponse
// @Failure      400  {object} response.Error
// @Failure      500  {object} response.Error
// @Router       /auth/register [post]
func (h *Handler) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req RegisterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			response.HandleError(c, err)
			return
		}

		if err := validator.Validate(req); err != nil {
			response.HandleError(c, webErrValidationFailed)
			return
		}

		input := ctrlAuth.RegisterInput{
			Name:     req.Name,
			Email:    req.Email,
			Password: req.Password,
		}

		token, err := h.ctrl.Register(c.Request.Context(), input)
		if err != nil {
			response.HandleError(c, err)
			return
		}

		c.JSON(http.StatusOK, RegisterResponse{Token: token})
	}
}
