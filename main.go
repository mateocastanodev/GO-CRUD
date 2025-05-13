package main

import (
	"example/crud/auth"
	"example/crud/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin router
	router := gin.Default()

	// Initialize auth service and handler
	authService := auth.NewService()
	authHandler := auth.NewHandler(authService)

	// Setup routes
	routes.SetupAuthRoutes(router, authHandler)

	// Start server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
