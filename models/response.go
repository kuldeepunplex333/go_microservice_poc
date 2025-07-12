package models

import "time"

// SimpleResponse represents a basic text response
type SimpleResponse struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

// NewSimpleResponse creates a new simple response
func NewSimpleResponse(message string) *SimpleResponse {
	return &SimpleResponse{
		Message:   message,
		Timestamp: time.Now(),
	}
}

// HealthResponse represents a health check response
type HealthResponse struct {
	Status    string    `json:"status"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

// NewHealthResponse creates a new health response
func NewHealthResponse(status, message string) *HealthResponse {
	return &HealthResponse{
		Status:    status,
		Message:   message,
		Timestamp: time.Now(),
	}
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error     string    `json:"error"`
	Code      int       `json:"code"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

// NewErrorResponse creates a new error response
func NewErrorResponse(code int, error, message string) *ErrorResponse {
	return &ErrorResponse{
		Error:     error,
		Code:      code,
		Message:   message,
		Timestamp: time.Now(),
	}
}
