package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/namf2001/beta-workplace/internal/handler/response"
	"github.com/namf2001/beta-workplace/internal/pkg/jwt"
)

var (
	webErrMissingAuth  = &response.Error{Status: http.StatusUnauthorized, Code: "missing_auth", Desc: "Missing authorization header"}
	webErrInvalidAuth  = &response.Error{Status: http.StatusUnauthorized, Code: "invalid_auth", Desc: "Invalid authorization header format"}
	webErrInvalidToken = &response.Error{Status: http.StatusUnauthorized, Code: "invalid_token", Desc: "Invalid or expired token"}
)

// RequireAuth middleware verifies JWT token
func RequireAuth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, webErrMissingAuth)
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, webErrInvalidAuth)
		return
	}

	tokenString := headerParts[1]
	claims, err := jwt.ParseToken(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, webErrInvalidToken)
		return
	}

	// Add UserID to context
	c.Set("userID", claims.UserID)
	c.Next()
}
