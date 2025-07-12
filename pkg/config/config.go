package config

import (
        "os"
        "strings"
)

// Config holds application configuration
type Config struct {
        Port        string
        Environment string
        LogLevel    string
}

// Load loads configuration from environment variables with defaults
func Load() *Config {
        return &Config{
                Port:        getEnv("PORT", "5000"),
                Environment: getEnv("ENVIRONMENT", "development"),
                LogLevel:    getEnv("LOG_LEVEL", "INFO"),
        }
}

// getEnv gets an environment variable with a default value
func getEnv(key, defaultValue string) string {
        if value := os.Getenv(key); value != "" {
                return value
        }
        return defaultValue
}

// IsProduction checks if the application is running in production
func (c *Config) IsProduction() bool {
        return strings.ToLower(c.Environment) == "production"
}

// IsDevelopment checks if the application is running in development
func (c *Config) IsDevelopment() bool {
        return strings.ToLower(c.Environment) == "development"
}
