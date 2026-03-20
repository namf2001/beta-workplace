package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/namf2001/beta-workplace/docs/swagger"
	appMiddleware "github.com/namf2001/beta-workplace/internal/handler/middleware"
	authhandler "github.com/namf2001/beta-workplace/internal/handler/rest/v1/auth"
	usershandler "github.com/namf2001/beta-workplace/internal/handler/rest/v1/users"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// router defines the routes & handlers of the app
type router struct {
	ctx          context.Context
	usersHandler *usershandler.Handler
	authHandler  *authhandler.Handler
}

// handler returns the handler for use by the server
func (rtr router) handler() http.Handler {
	r := gin.New()

	// Middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300 * time.Second,
	}))

	rtr.routes(r)

	return r
}

func (rtr router) routes(r *gin.Engine) {
	rtr.public(r)
	rtr.apiV1(r)
}

func (rtr router) public(r *gin.Engine) {
	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	// Metrics
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (rtr router) apiV1(r *gin.Engine) {
	v1 := r.Group("/api/v1")

	auth := v1.Group("/auth")
	{
		auth.POST("/login", rtr.authHandler.Login())
		auth.POST("/register", rtr.authHandler.Register())
		auth.POST("/forgot-password", rtr.authHandler.ForgotPassword())
		auth.GET("/google/login", rtr.authHandler.GoogleLogin())
		auth.GET("/google/callback", rtr.authHandler.GoogleCallback())
	}

	protected := v1.Group("/")
	protected.Use(appMiddleware.RequireAuth)
	{
		authProtected := protected.Group("/auth")
		{
			authProtected.POST("/logout", rtr.authHandler.Logout())

			// User Management APIs
			userManagement := authProtected.Group("/user")
			{
				userManagement.GET("/profile", rtr.authHandler.GetProfile())
				userManagement.PUT("/profile", rtr.authHandler.UpdateProfile())
				userManagement.PATCH("/password", rtr.authHandler.ChangePassword())
				userManagement.DELETE("/account", rtr.authHandler.DeleteAccount())
			}
		}

		users := protected.Group("/users")
		{
			users.POST("", rtr.usersHandler.CreateUser())
			users.GET("", rtr.usersHandler.ListUsers())
			users.GET("/:id", rtr.usersHandler.GetUser())
			users.PUT("/:id", rtr.usersHandler.UpdateUser())
			users.DELETE("/:id", rtr.usersHandler.DeleteUser())
		}
	}
}
