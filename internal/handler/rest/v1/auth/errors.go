package auth

import (
	"net/http"

	"github.com/namf2001/beta-workplace/internal/handler/response"
)

var (
	webErrValidationFailed   = &response.Error{Status: http.StatusBadRequest, Code: "validation_failed", Desc: "Validation failed"}
	webErrInvalidCredentials = &response.Error{Status: http.StatusUnauthorized, Code: "invalid_credentials", Desc: "Invalid email or password"}
	webErrInvalidOAuthState  = &response.Error{Status: http.StatusBadRequest, Code: "invalid_oauth_state", Desc: "Invalid OAuth state"}
	webErrCodeExchangeFailed = &response.Error{Status: http.StatusInternalServerError, Code: "code_exchange_failed", Desc: "Failed to exchange code"}
	webErrGetUserInfoFailed  = &response.Error{Status: http.StatusInternalServerError, Code: "get_user_info_failed", Desc: "Failed to get user info"}
)
