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

type ForgotPasswordRequest struct {
	Step        int    `json:"step" validate:"required,oneof=1 2"`
	Email       string `json:"email" validate:"required,email"`
	OTP         string `json:"otp"`          // required for step 2
	NewPassword string `json:"new_password"` // required for step 2
}

// ForgotPassword handles multi-step forgot password flow
// @Summary      Forgot Password
// @Description  Forgot password via 2-step OTP flow
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input body auth.ForgotPasswordRequest true "Forgot password info"
// @Success      200  {object} response.Response
// @Failure      400  {object} response.Response
// @Failure      500  {object} response.Response
// @Router       /auth/forgot-password [post]
func (h Handler) ForgotPassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req ForgotPasswordRequest
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
			err := h.ctrl.ForgotPasswordStep1SendOTP(ctx, req.Email)
			if err != nil {
				code := constants.InternalServerError.Code
				if errors.Is(err, ctrlAuth.ErrUserNotFound) {
					code = constants.UserNotFound.Code
				}
				c.JSON(http.StatusInternalServerError, response.NewResponse(
					code,
					err.Error(),
					nil,
				))
				return
			}
			c.JSON(http.StatusOK, response.NewResponse(
				constants.SendMailForgotPasswordSuccess.Code,
				constants.SendMailForgotPasswordSuccess.Message,
				nil,
			))
			return

		case 2:
			if req.OTP == "" || len(req.NewPassword) < 6 {
				c.JSON(http.StatusBadRequest, response.NewResponse(
					constants.BindJSONFail.Code,
					constants.BindJSONFail.Message,
					nil,
				))
				return
			}

			err := h.ctrl.ForgotPasswordStep2Reset(ctx, req.Email, req.OTP, req.NewPassword)
			if err != nil {
				code := constants.InternalServerError.Code
				if errors.Is(err, ctrlAuth.ErrWrongOTP) {
					code = constants.VerifyCodeExpired.Code
				} else if errors.Is(err, ctrlAuth.ErrUserNotFound) {
					code = constants.UserNotFound.Code
				}
				c.JSON(http.StatusInternalServerError, response.NewResponse(
					code,
					err.Error(),
					nil,
				))
				return
			}
			c.JSON(http.StatusOK, response.NewResponse(
				constants.ChangePasswordSuccess.Code,
				constants.ChangePasswordSuccess.Message,
				nil,
			))
			return
		}
	}
}
