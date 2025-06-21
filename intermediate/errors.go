package main

// This file demonstrates error handling in Go, which uses explicit error values
// instead of exceptions. Go functions often return an error value to indicate
// success or failure, and the caller is responsible for checking and handling errors.

import "errors" // Package errors implements functions to manipulate errors

// sqrt calculates the square of a number and demonstrates error handling
// Input: x - the number to calculate the square of
// Returns:
//   - float64: the square of x if x is non-negative
//   - error: nil if successful, otherwise an error explaining why it failed
//
// Note: This function actually calculates square, not square root despite its name
func sqrt(x float64) (float64, error) {
	// Check if the input is valid
	if x < 0 {
		// If not valid, return zero value and an error
		// errors.New creates a basic error with the given text
		return 0, errors.New("cannot compute square root of negative number")
	}
	// If valid, return the result and nil error
	// nil indicates that no error occurred
	return x * x, nil

	// Note: In real code, you would likely:
	// 1. Use a more specific error type
	// 2. Use fmt.Errorf for formatted errors: fmt.Errorf("cannot compute square root of %v", x)
	// 3. Consider using error wrapping: fmt.Errorf("calculation failed: %w", err)
}

func main() {
	// Example 1: Call sqrt with a negative number (should result in an error)
	result, err := sqrt(-4)

	// Handle the error using Go's idiomatic error checking pattern
	// By convention, we check err != nil to determine if an error occurred
	if err != nil {
		// If an error occurred, handle it appropriately
		// Here we just print the error message
		println("Error:", err.Error())
	} else {
		// If no error (err == nil), use the result
		println("Square root:", result)
	}

	// Example 2: Call sqrt with a positive number (should succeed)
	result, err = sqrt(16)
	if err != nil {
		println("Error:", err.Error())
	} else {
		println("Square root:", result) // Will print 256 (16²)
	}

	// Note: There are several other error handling techniques in Go:
	// 1. Error wrapping and unwrapping (Go 1.13+)
	// 2. Custom error types (implementing the error interface)
	// 3. Error inspection with errors.Is and errors.As (Go 1.13+)
	// 4. Sentinel errors (predefined error values like io.EOF)
}
