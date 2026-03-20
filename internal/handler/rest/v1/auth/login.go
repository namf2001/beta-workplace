package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/namf2001/beta-workplace/constants"
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
// @Success      200  {object} response.Response
// @Failure      400  {object} response.Response
// @Failure      401  {object} response.Response
// @Failure      500  {object} response.Response
// @Router       /auth/login [post]
func (h Handler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LoginRequest
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

		input := ctrlAuth.ValidationInput{
			Email:    req.Email,
			Password: req.Password,
		}

		token, err := h.ctrl.Login(c.Request.Context(), input)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.NewResponse(
				constants.LoginFail.Code,
				err.Error(),
				nil,
			))
			return
		}

		c.JSON(http.StatusOK, response.NewResponse(
			constants.LoginSuccess.Code,
			constants.LoginSuccess.Message,
			LoginResponse{Token: token},
		))
	}
}
