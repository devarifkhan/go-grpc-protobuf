package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// line filters: programs that read input on stdin and print output on stdout
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Scan input line by line
	for scanner.Scan() {
		line := scanner.Text()

		// Process the line (for example, convert to uppercase)
		processedLine := strings.ToUpper(line)

		// Print the processed line
		fmt.Println(processedLine)
	}

	// Check for errors
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
