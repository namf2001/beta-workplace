package auth

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/namf2001/beta-workplace/constants"
	ctrlAuth "github.com/namf2001/beta-workplace/internal/controller/auth"
	"github.com/namf2001/beta-workplace/internal/handler/response"
	"github.com/namf2001/beta-workplace/internal/model"
	"github.com/namf2001/beta-workplace/internal/pkg/oauth"
)

const (
	GoogleUserInfoURL = "https://www.googleapis.com/oauth2/v2/userinfo"
)

// GoogleLoginResponse represents the response for Google login
type GoogleLoginResponse struct {
	URL string `json:"url"`
}

// GoogleCallbackResponse represents the response for Google callback
type GoogleCallbackResponse struct {
	Token string `json:"token"`
}

type GoogleUserInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
}

// GoogleLogin handles google login
// @Summary      Google login
// @Description  Get Google login URL
// @Tags         auth
// @Produce      json
// @Success      200  {object} auth.GoogleLoginResponse
// @Router       /auth/google/login [get]
func (h Handler) GoogleLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := oauth.GoogleOauthConfig.AuthCodeURL(oauth.StateString)
		c.JSON(http.StatusOK, response.NewResponse(
			constants.LoginSuccess.Code,
			constants.LoginSuccess.Message,
			GoogleLoginResponse{URL: url},
		))
	}
}

// GoogleCallback handles google callback
// @Summary      Google callback
// @Description  Handle Google OAuth callback and return token
// @Tags         auth
// @Produce      json
// @Param        state query string true "OAuth state"
// @Param        code  query string true "OAuth code"
// @Success      200  {object} auth.GoogleCallbackResponse
// @Failure      400  {object} response.Response
// @Failure      500  {object} response.Response
// @Router       /auth/google/callback [get]
func (h Handler) GoogleCallback() gin.HandlerFunc {
	return func(c *gin.Context) {
		state := c.Query("state")
		if state != oauth.StateString {
			c.JSON(http.StatusBadRequest, response.NewResponse(
				constants.InvalidOAuthState.Code,
				constants.InvalidOAuthState.Message,
				nil,
			))
			return
		}

		code := c.Query("code")
		token, err := oauth.GoogleOauthConfig.Exchange(c.Request.Context(), code)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.NewResponse(
				constants.CodeExchangeFailed.Code,
				err.Error(),
				nil,
			))
			return
		}

		resp, err := http.Get(GoogleUserInfoURL + "?access_token=" + token.AccessToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.NewResponse(
				constants.GetUserInfoFail.Code,
				err.Error(),
				nil,
			))
			return
		}
		defer func() { _ = resp.Body.Close() }()

		content, err := io.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.NewResponse(
				constants.GetUserInfoFail.Code,
				err.Error(),
				nil,
			))
			return
		}

		var userInfo GoogleUserInfo
		if err := json.Unmarshal(content, &userInfo); err != nil {
			c.JSON(http.StatusInternalServerError, response.NewResponse(
				constants.GetUserInfoFail.Code,
				err.Error(),
				nil,
			))
			return
		}

		input := ctrlAuth.OAuthInput{
			Provider:          model.ProviderGoogle,
			ProviderAccountID: userInfo.ID,
			Type:              "oauth",
			AccessToken:       token.AccessToken,
			RefreshToken:      token.RefreshToken,
			ExpiresAt:         token.Expiry.Unix(),
			TokenType:         token.TokenType,
			Name:              userInfo.Name,
			Email:             userInfo.Email,
			Image:             userInfo.Picture,
			EmailVerified:     userInfo.VerifiedEmail,
		}

		authToken, err := h.ctrl.OAuthLogin(c.Request.Context(), input)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.NewResponse(
				constants.InternalServerError.Code,
				err.Error(),
				nil,
			))
			return
		}

		c.JSON(http.StatusOK, response.NewResponse(
			constants.LoginSuccess.Code,
			constants.LoginSuccess.Message,
			GoogleCallbackResponse{Token: authToken},
		))
	}
}
