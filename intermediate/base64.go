package main

// This file demonstrates how to use base64 encoding in Go.
// Base64 encoding is commonly used to encode binary data into ASCII strings,
// making it easier to transmit data across text-based protocols like HTTP.

import (
	"encoding/base64"
	"fmt"
	"io/ioutil" // Note: ioutil is deprecated in Go 1.16+, use io and os packages instead
)

func main() {
	// Read the entire contents of a file into a byte slice
	// This provides the binary data we'll encode to base64
	data, err := ioutil.ReadFile("bufio.txt")
	if err != nil {
		// Handle any errors that occurred during file reading
		fmt.Println(err)
		return
	}

	// Encode the binary data to a base64 string
	// StdEncoding uses the standard base64 encoding with padding
	// There's also URLEncoding for URL-safe base64, and RawStdEncoding/RawURLEncoding for unpadded variants
	encoded := base64.StdEncoding.EncodeToString(data)

	// Print the base64 encoded string
	fmt.Println(encoded)

	// Note: To decode base64 back to binary:
	// decoded, err := base64.StdEncoding.DecodeString(encoded)
}
