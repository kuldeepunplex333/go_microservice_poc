package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-microservice/models"
	"go-microservice/pkg/logger"
)

// APIController handles API-related requests
type APIController struct{}

// NewAPIController creates a new API controller instance
func NewAPIController() *APIController {
	return &APIController{}
}

// HelloWorld handles the GET /api endpoint
// Returns a simple "hello" message in plain text
func (ac *APIController) HelloWorld(c *gin.Context) {
	logger.Info("HelloWorld endpoint called from IP: %s", c.ClientIP())
	
	// Create response model
	response := models.NewSimpleResponse("hello")
	
	// Log the response
	logger.Debug("Returning response: %s", response.Message)
	
	// Return plain text response
	c.String(http.StatusOK, response.Message)
}

// HealthCheck handles health check requests
func (ac *APIController) HealthCheck(c *gin.Context) {
	logger.Debug("Health check endpoint called")
	
	response := models.NewHealthResponse("healthy", "Service is running")
	
	c.JSON(http.StatusOK, response)
}
