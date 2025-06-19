package main

import (
	"fmt"
	"net/url"
)

// URLParsingExample demonstrates how to parse URLs in Go
func main() {

	// Example URL
	exampleURL := "https://www.google.com/search?q=golang#section1"

	// Parse the URL
	parsedURL, err := url.Parse(exampleURL)
	if err != nil {
		panic(err)
	}

	// Print the parsed components
	fmt.Println("Scheme:", parsedURL.Scheme)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("Path:", parsedURL.Path)
	fmt.Println("Query:", parsedURL.RawQuery)
	fmt.Println("Fragment:", parsedURL.Fragment)

}
