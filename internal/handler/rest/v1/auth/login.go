package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ctrlAuth "github.com/namf2001/beta-workplace/internal/controller/auth"
	"github.com/namf2001/beta-workplace/internal/handler/response"
	"github.com/namf2001/beta-workplace/internal/pkg/validator"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// Login handles manual login
// @Summary      User login
// @Description  Authenticate user and return token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input body auth.LoginRequest true "Login credentials"
// @Success      200  {object} auth.LoginResponse
// @Failure      400  {object} response.Error
// @Failure      401  {object} response.Error
// @Failure      500  {object} response.Error
// @Router       /auth/login [post]
func (h *Handler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			response.HandleError(c, err)
			return
		}

		if err := validator.Validate(req); err != nil {
			response.HandleError(c, webErrValidationFailed)
			return
		}

		input := ctrlAuth.ValidationInput{
			Email:    req.Email,
			Password: req.Password,
		}

		token, err := h.ctrl.Login(c.Request.Context(), input)
		if err != nil {
			response.HandleError(c, webErrInvalidCredentials)
			return
		}

		c.JSON(http.StatusOK, LoginResponse{Token: token})
	}
}
