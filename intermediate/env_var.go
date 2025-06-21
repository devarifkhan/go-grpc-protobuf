package main

// This file demonstrates how to work with environment variables in Go.
// Environment variables are key-value pairs used to configure applications
// in different environments without changing the code. They're commonly used
// for configuration, secrets, and environment-specific settings.

import (
	"fmt"     // Package fmt implements formatted I/O
	"os"      // Package os provides platform-independent OS functionality
	"strconv" // Package strconv implements string conversions to various types
)

func main() {
	// 1. Setting environment variables programmatically
	// os.Setenv sets the value of the environment variable identified by key
	// Note: These changes are only visible to the current process and its children
	os.Setenv("APP_NAME", "Go GRPC Service")
	os.Setenv("APP_PORT", "8080")
	os.Setenv("DEBUG_MODE", "true")
	os.Setenv("TIMEOUT_SECONDS", "30")

	// 2. Reading environment variables directly
	// os.Getenv retrieves the value of the environment variable named by the key
	// If the variable is not present, it returns an empty string
	appName := os.Getenv("APP_NAME")
	fmt.Println("App Name:", appName)

	// 3. Reading with a default value if the environment variable isn't set
	// This uses our helper function defined below
	appPort := getEnvOrDefault("APP_PORT", "9000")
	fmt.Println("Port:", appPort)

	// 4. Converting environment variables to appropriate types
	// Environment variables are strings, so we need to convert them to the desired type
	// The _ ignores the error return value (not recommended for production code)
	debugMode, _ := strconv.ParseBool(os.Getenv("DEBUG_MODE"))
	fmt.Println("Debug Mode:", debugMode)

	// 5. Using a helper function for getting and parsing integer variables
	timeout := getEnvAsInt("TIMEOUT_SECONDS", 60)
	fmt.Println("Timeout:", timeout, "seconds")

	// 6. Getting all environment variables
	// os.Environ returns a copy of strings representing the environment,
	// in the form "key=value"
	fmt.Println("\nAll Environment Variables:")
	for _, env := range os.Environ() {
		fmt.Println(env)
	}

	// Note: For production applications, consider using a dedicated configuration
	// package like github.com/spf13/viper or github.com/kelseyhightower/envconfig
}

// getEnvOrDefault retrieves an environment variable with a fallback default value
// key: the name of the environment variable
// defaultValue: the value to use if the environment variable is not set
// Returns: the environment variable value or the default if not set
func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		// Return the default value if the environment variable is not set
		// or is an empty string
		return defaultValue
	}
	return value
}

// getEnvAsInt retrieves an environment variable and converts it to an integer
// key: the name of the environment variable
// defaultValue: the integer to use if the variable is not set or cannot be parsed
// Returns: the parsed integer value or the default value
func getEnvAsInt(key string, defaultValue int) int {
	// Get the environment variable as a string
	valueStr := os.Getenv(key)
	if valueStr == "" {
		// Return the default if not set
		return defaultValue
	}

	// Try to parse the string as an integer
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		// If parsing fails (e.g., not a valid integer), return the default
		return defaultValue
	}

	// Return the successfully parsed integer
	return value

	// Note: You could create similar helpers for other types:
	// - getEnvAsBool
	// - getEnvAsFloat
	// - getEnvAsTime
	// etc.
}
