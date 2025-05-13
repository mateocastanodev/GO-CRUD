package auth

import (
	"github.com/gin-gonic/gin"
)

// Middleware represents authentication middleware
type Middleware struct {
	service *Service
}

// NewMiddleware creates a new auth middleware
func NewMiddleware(service *Service) *Middleware {
	return &Middleware{
		service: service,
	}
}

// AuthRequired is a placeholder middleware that will be implemented later
func (m *Middleware) AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// This is a placeholder for future authentication middleware
		// It will be implemented when we add JWT or session-based auth
		c.Next()
	}
}
