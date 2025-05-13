package routes

import (
	"example/crud/auth"

	"github.com/gin-gonic/gin"
)

// SetupAuthRoutes configures the authentication routes
func SetupAuthRoutes(router *gin.Engine, authHandler *auth.Handler) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}
}
