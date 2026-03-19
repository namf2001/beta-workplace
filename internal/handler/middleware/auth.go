package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/namf2001/beta-workplace/constants"
	"github.com/namf2001/beta-workplace/internal/handler/response"
	"github.com/namf2001/beta-workplace/internal/pkg/jwt"
)

// RequireAuth middleware verifies JWT token
func RequireAuth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, response.NewResponse(
			constants.MissingAuthorizationHeader.Code,
			constants.MissingAuthorizationHeader.Message,
			nil,
		))
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, response.NewResponse(
			constants.InvalidAuthorizationHeader.Code,
			constants.InvalidAuthorizationHeader.Message,
			nil,
		))
		return
	}

	tokenString := headerParts[1]
	claims, err := jwt.ParseToken(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, response.NewResponse(
			constants.InvalidToken.Code,
			constants.InvalidToken.Message,
			nil,
		))
		return
	}

	// Add UserID to context
	c.Set("userID", claims.UserID)
	c.Next()
}
