package auth

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/namf2001/beta-workplace/constants"
	ctrlAuth "github.com/namf2001/beta-workplace/internal/controller/auth"
	"github.com/namf2001/beta-workplace/internal/handler/response"
	"github.com/namf2001/beta-workplace/internal/pkg/validator"
)

type RegisterRequest struct {
	Step     int    `json:"step" validate:"required,oneof=1 2 3"`
	Email    string `json:"email" validate:"required,email"`
	OTP      string `json:"otp"`      // required for step 2 & 3
	Name     string `json:"name"`     // required for step 3
	Password string `json:"password"` // required for step 3
}

type RegisterResponse struct {
	Token string `json:"token"`
}

// Register handles manual registration, multi-step OTP
// @Summary      Register user
// @Description  Register a new user account via 3-step OTP flow
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input body auth.RegisterRequest true "Registration info"
// @Success      200  {object} response.Response
// @Failure      400  {object} response.Response
// @Failure      500  {object} response.Response
// @Router       /auth/register [post]
func (h Handler) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req RegisterRequest
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

		ctx := c.Request.Context()

		switch req.Step {
		case 1:
			err := h.ctrl.RegisterStep1SendOTP(ctx, req.Email)
			if err != nil {
				code := constants.InternalServerError.Code
				if errors.Is(err, ctrlAuth.ErrEmailAlreadyExists) {
					code = constants.EmailExists.Code
				}
				c.JSON(http.StatusInternalServerError, response.NewResponse(
					code,
					err.Error(),
					nil,
				))
				return
			}
			c.JSON(http.StatusOK, response.NewResponse(
				constants.SendEmailRegisterSuccess.Code,
				constants.SendEmailRegisterSuccess.Message,
				nil,
			))
			return

		case 2:
			if req.OTP == "" {
				c.JSON(http.StatusBadRequest, response.NewResponse(
					constants.InvalidRequestParams.Code,
					constants.InvalidRequestParams.Message,
					nil,
				))
				return
			}
			err := h.ctrl.RegisterStep2VerifyOTP(ctx, req.Email, req.OTP)
			if err != nil {
				code := constants.InternalServerError.Code
				if errors.Is(err, ctrlAuth.ErrWrongOTP) {
					code = constants.VerifyCodeExpired.Code
				}
				c.JSON(http.StatusInternalServerError, response.NewResponse(
					code,
					err.Error(),
					nil,
				))
				return
			}
			c.JSON(http.StatusOK, response.NewResponse(
				constants.EmailVerifiedSuccess.Code,
				constants.EmailVerifiedSuccess.Message,
				nil,
			))
			return

		case 3:
			if req.OTP == "" || req.Name == "" || len(req.Password) < 6 {
				c.JSON(http.StatusBadRequest, response.NewResponse(
					constants.InvalidRequestParams.Code,
					constants.InvalidRequestParams.Message,
					nil,
				))
				return
			}

			input := ctrlAuth.RegisterInput{
				Name:     req.Name,
				Email:    req.Email,
				OTP:      req.OTP,
				Password: req.Password,
			}

			token, err := h.ctrl.RegisterStep3Complete(ctx, input)
			if err != nil {
				code := constants.InternalServerError.Code
				if errors.Is(err, ctrlAuth.ErrEmailAlreadyExists) {
					code = constants.EmailExists.Code
				} else if errors.Is(err, ctrlAuth.ErrWrongOTP) {
					code = constants.VerifyCodeExpired.Code
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
				RegisterResponse{Token: token},
			))
			return
		}
	}
}
