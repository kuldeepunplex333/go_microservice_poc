package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// LogLevel represents the logging level
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

var (
	currentLevel LogLevel = INFO
	logger       *log.Logger
)

// Init initializes the logger with the specified level
func Init(level string) {
	logger = log.New(os.Stdout, "", 0)
	
	switch strings.ToUpper(level) {
	case "DEBUG":
		currentLevel = DEBUG
	case "INFO":
		currentLevel = INFO
	case "WARN":
		currentLevel = WARN
	case "ERROR":
		currentLevel = ERROR
	default:
		currentLevel = INFO
	}
}

// formatMessage creates a formatted log message
func formatMessage(level string, format string, args ...interface{}) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	message := fmt.Sprintf(format, args...)
	return fmt.Sprintf("[%s] %s: %s", timestamp, level, message)
}

// Debug logs a debug message
func Debug(format string, args ...interface{}) {
	if currentLevel <= DEBUG {
		logger.Println(formatMessage("DEBUG", format, args...))
	}
}

// Info logs an info message
func Info(format string, args ...interface{}) {
	if currentLevel <= INFO {
		logger.Println(formatMessage("INFO", format, args...))
	}
}

// Warn logs a warning message
func Warn(format string, args ...interface{}) {
	if currentLevel <= WARN {
		logger.Println(formatMessage("WARN", format, args...))
	}
}

// Error logs an error message
func Error(format string, args ...interface{}) {
	if currentLevel <= ERROR {
		logger.Println(formatMessage("ERROR", format, args...))
	}
}

// Fatal logs a fatal message and exits
func Fatal(format string, args ...interface{}) {
	logger.Println(formatMessage("FATAL", format, args...))
	os.Exit(1)
}
