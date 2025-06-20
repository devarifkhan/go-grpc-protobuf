package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Setting environment variables
	os.Setenv("APP_NAME", "Go GRPC Service")
	os.Setenv("APP_PORT", "8080")
	os.Setenv("DEBUG_MODE", "true")
	os.Setenv("TIMEOUT_SECONDS", "30")

	// Getting environment variables
	appName := os.Getenv("APP_NAME")
	fmt.Println("App Name:", appName)

	// Getting with default value if not set
	appPort := getEnvOrDefault("APP_PORT", "9000")
	fmt.Println("Port:", appPort)

	// Converting to appropriate types
	debugMode, _ := strconv.ParseBool(os.Getenv("DEBUG_MODE"))
	fmt.Println("Debug Mode:", debugMode)

	// Getting and parsing with helper function
	timeout := getEnvAsInt("TIMEOUT_SECONDS", 60)
	fmt.Println("Timeout:", timeout, "seconds")

	// Getting all environment variables
	fmt.Println("\nAll Environment Variables:")
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
}

// Helper function to get env var with default value
func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// Helper function to get env var as integer
func getEnvAsInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}
