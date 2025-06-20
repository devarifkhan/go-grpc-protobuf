package main

import (
	"log"
	"os"
	"time"
)

// Logger represents a simple logger for the application
type Logger struct {
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
}

// NewLogger creates and returns a new Logger instance
func NewLogger() *Logger {
	return &Logger{
		infoLogger:    log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime),
		warningLogger: log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime),
		errorLogger:   log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// Info logs information messages
func (l *Logger) Info(format string, v ...interface{}) {
	l.infoLogger.Printf(format, v...)
}

// Warning logs warning messages
func (l *Logger) Warning(format string, v ...interface{}) {
	l.warningLogger.Printf(format, v...)
}

// Error logs error messages
func (l *Logger) Error(format string, v ...interface{}) {
	l.errorLogger.Printf(format, v...)
}

// LogDuration logs the duration of a function execution
func (l *Logger) LogDuration(name string) func() {
	start := time.Now()
	return func() {
		l.Info("%s took %v", name, time.Since(start))
	}
}
