package main

// This file demonstrates how to use the bufio package for buffered I/O operations.
// Buffered I/O operations reduce the number of system calls by batching reads and writes,
// providing more efficient input/output operations, especially for line-by-line reading.

import (
	"bufio" // Package bufio implements buffered I/O
	"fmt"   // Package fmt implements formatted I/O
	"os"    // Package os provides platform-independent OS functionality
)

func main() {
	// Open a file for reading
	// os.Open returns a file handle (*os.File) and an error, if any
	file, err := os.Open("bufio.txt")
	if err != nil {
		// Handle any errors during file opening (e.g., file not found)
		fmt.Println(err)
		return
	}
	// Ensure the file is closed when the function exits
	// defer executes the statement when the surrounding function returns
	defer file.Close()

	// Create a new Scanner from the file
	// Scanner provides a convenient way to read data line by line
	scanner := bufio.NewScanner(file)

	// Scan advances the scanner to the next token (line in this case)
	// and returns true if there's a token, false if at EOF or an error
	for scanner.Scan() {
		// Text returns the current token (line) as a string
		line := scanner.Text()
		fmt.Println(line)

		// You can also use scanner.Bytes() to get the line as a []byte
	}

	// Check if an error occurred during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	// Note: bufio also provides Reader and Writer types for more
	// fine-grained buffered I/O operations
}
