package main

import (
        "log"
        "net/http"

        "github.com/gin-gonic/gin"
        "go-microservice/middleware"
        "go-microservice/pkg/config"
        "go-microservice/pkg/logger"
        "go-microservice/routes"
)

func main() {
        // Initialize configuration
        cfg := config.Load()
        
        // Initialize logger
        logger.Init(cfg.LogLevel)
        
        // Set Gin mode based on environment
        if cfg.Environment == "production" {
                gin.SetMode(gin.ReleaseMode)
        }
        
        // Initialize Gin router
        router := gin.New()
        
        // Add middleware
        router.Use(gin.Recovery())
        router.Use(middleware.LoggerMiddleware())
        router.Use(middleware.CORSMiddleware())
        
        // Setup routes
        routes.SetupRoutes(router)
        
        // Start server
        logger.Info("Starting server on port %s", cfg.Port)
        logger.Info("Environment: %s", cfg.Environment)
        
        server := &http.Server{
                Addr:    "0.0.0.0:" + cfg.Port,
                Handler: router,
        }
        // main
        // second commit 
        // third commit
        
        
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
                log.Fatalf("Failed to start server: %v", err)
        }
}
