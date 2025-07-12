package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go-microservice/pkg/logger"
)

// LoggerMiddleware creates a middleware for logging HTTP requests
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		
		// Process request
		c.Next()
		
		// Calculate latency
		latency := time.Since(start)
		
		// Get request details
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		
		// Build path with query parameters
		if raw != "" {
			path = path + "?" + raw
		}
		
		// Log the request
		logger.Info("%s %s %d %v %s",
			method,
			path,
			statusCode,
			latency,
			clientIP,
		)
		
		// Log errors if any
		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				logger.Error("Request error: %s", err.Error())
			}
		}
	}
}
