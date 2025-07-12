# Go Microservice Project Documentation

## Overview

This is a Go microservice built with the Gin web framework following clean architecture principles. The project implements a simple "Hello World" API with proper logging, middleware, and configuration management.

## User Preferences

Preferred communication style: Simple, everyday language.

## System Architecture

### Backend Architecture
- Runtime environment: Go 1.19
- Framework: Gin web framework
- API design: RESTful API with JSON/text responses
- Port: 5000 (Replit-compatible)

### Folder Structure
```
/controllers     // Business logic handlers
/middleware      // CORS, logging middleware
/models          // Data models and response structures
/pkg             // Shared utilities (config, logger)
/routes          // Route definitions
main.go          // Application entry point
go.mod           // Go module dependencies
```

## Key Components

### Core Modules
- Main application entry point: `main.go` - bootstraps Gin, registers routes, starts server
- Configuration management: `pkg/config/config.go` - environment-based configuration
- Error handling: Built into controllers and middleware

### Business Logic
- Domain models: Response structures in `/models`
- Service layers: Controllers in `/controllers`
- Validation logic: Gin framework validation

### API Endpoints
- `GET /api` - Returns "hello" in plain text
- `GET /api/health` - Health check endpoint (JSON response)
- `GET /health` - Root health check endpoint

## Data Flow

### Request/Response Cycle
- Request handling: Gin router with middleware stack
- Data processing: Controllers process requests and return responses
- Response formatting: Plain text for hello endpoint, JSON for health checks

### Middleware Stack
1. Gin Recovery middleware
2. Custom logging middleware
3. CORS middleware
4. Route handlers

## External Dependencies

### Package Dependencies
- Core dependencies: github.com/gin-gonic/gin v1.9.1
- Development dependencies: Standard Go toolchain
- Version management: Go modules (go.mod)

## Deployment Strategy

### Environment Configuration
- Development setup: Local development on port 5000
- Production requirements: Go 1.19+, configurable via environment variables
- Environment variables: PORT (default: 5000), ENVIRONMENT (default: development), LOG_LEVEL (default: INFO)

### Build and Deploy
- Build process: `go mod tidy && go run main.go`
- Deployment target: Replit platform
- CI/CD pipeline: Replit workflows

## Recent Changes

### July 12, 2025
- ✓ Created complete Go microservice scaffold with Gin framework
- ✓ Implemented proper folder structure with controllers, middleware, models, pkg, routes
- ✓ Added "Hello World" endpoint at GET /api returning plain text "hello"
- ✓ Implemented CORS and logging middleware
- ✓ Added health check endpoints for monitoring
- ✓ Fixed Go version compatibility (1.21 → 1.19)
- ✓ Configured server to run on port 5000 for Replit compatibility
- ✓ Successfully tested endpoint functionality