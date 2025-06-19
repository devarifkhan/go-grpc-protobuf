package main

// string and runes in go
func main() {

	// Example usage of strings and runes
	str := "Hello, GO" // A string containing both ASCII and non-ASCII characters
	println("Original string:", str)

	// Convert string to rune slice
	runes := []rune(str)
	println("Number of runes:", len(runes))

	// Iterate over runes
	for i, r := range runes {
		println("Rune", i, ":", string(r), "Unicode:", r)
	}

	// Convert rune slice back to string
	newStr := string(runes)
	println("Converted back to string:", newStr)
}
