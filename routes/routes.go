package routes

import (
	"github.com/gin-gonic/gin"
	"go-microservice/controllers"
)

// SetupRoutes configures all application routes
func SetupRoutes(router *gin.Engine) {
	// Initialize controllers
	apiController := controllers.NewAPIController()
	
	// API routes group
	api := router.Group("/api")
	{
		// Main hello world endpoint
		api.GET("", apiController.HelloWorld)
		
		// Health check endpoint
		api.GET("/health", apiController.HealthCheck)
	}
	
	// Root health check
	router.GET("/health", apiController.HealthCheck)
	
	// Handle 404 for undefined routes
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"error":   "not_found",
			"message": "The requested endpoint does not exist",
			"path":    c.Request.URL.Path,
		})
	})
}
