package main

import "fmt"

// formatting_verbs
func main() {

	// Example usage of formatting verbs
	name := "Alice"
	age := 30
	// Using Printf to format strings
	fmt.Printf("Name: %s, Age: %d\n", name, age) // Output: Name: Alice, Age: 30
	// Using Printf to format strings
	// Using Sprintf to create a formatted string
	formattedString := fmt.Sprintf("Hello, my name is %s and I am %d years old.", name, age)
	println(formattedString) // Output: Hello, my name is Alice and I am 30 years old.
}
