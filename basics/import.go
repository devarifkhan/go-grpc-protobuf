// Definition: The import statement is used to include packages in a Go program.
// How it works in Go: Uses the import keyword to bring in standard or third-party packages.
// Purpose: To show how to import and use packages in Go.

package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Hello, Go Standard Library!")

	// Make a simple HTTP GET request
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Printf("Error fetching data: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Success!")
	fmt.Println("Response status:", resp.Status)

	// Read and display response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Printf("Response body:\n%s\n", string(body))
}
