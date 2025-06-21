package main

// This file demonstrates how to access and process command-line arguments in Go.
// Command-line arguments are values passed to a program when it is executed from
// a command-line interface (CLI). They allow users to customize program behavior.

import (
	"fmt" // Package fmt implements formatted I/O
	"os"  // Package os provides platform-independent OS functionality, including args access
)

func main() {
	// os.Args provides access to raw command-line arguments as a slice of strings
	// os.Args[0] is always the program name (the path to the executable)
	// os.Args[1:] contains the actual arguments passed to the program

	// Display the program name (the executable path)
	fmt.Println("Program name:", os.Args[0])

	// Check if any command-line arguments were provided
	if len(os.Args) > 1 {
		fmt.Println("Arguments:")
		// Iterate through each argument using range
		// We use os.Args[1:] to skip the program name
		for i, arg := range os.Args[1:] {
			// Print each argument with its position (1-based for user-friendliness)
			fmt.Printf("  %d: %s\n", i+1, arg)
		}
	} else {
		fmt.Println("No command line arguments provided")
	}

	// Example usage:
	// Run with: go run command_line_arguments.go arg1 arg2 "arg with spaces"

	// For more advanced command-line argument parsing, consider using:
	// - The flag package from the standard library
	// - Third-party packages like github.com/spf13/cobra or github.com/urfave/cli
}
