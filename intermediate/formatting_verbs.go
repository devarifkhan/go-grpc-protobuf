package main

// This file demonstrates formatting verbs in Go using the fmt package.
// Formatting verbs are placeholders in format strings that are replaced
// with values when you use functions like Printf, Sprintf, etc.
// They allow you to control how values are formatted in output strings.

import "fmt" // Package fmt implements formatted I/O with functions similar to C's printf
func main() {
	// Initialize some variables to use in our formatting examples
	name := "Alice"
	age := 30
	pi := 3.14159
	isStudent := true
	data := []int{1, 2, 3}

	// Common formatting verbs:

	// %v - Default format (works with any value type)
	fmt.Printf("Default: %v %v %v %v %v\n", name, age, pi, isStudent, data)

	// %T - Type of the value
	fmt.Printf("Types: %T %T %T %T %T\n", name, age, pi, isStudent, data)

	// String formatting
	fmt.Printf("String: %s\n", name)        // %s - String
	fmt.Printf("Quoted string: %q\n", name) // %q - Quoted string

	// Integer formatting
	fmt.Printf("Integer: %d\n", age)     // %d - Decimal integer
	fmt.Printf("Binary: %b\n", age)      // %b - Binary
	fmt.Printf("Hex: %x %X\n", age, age) // %x/%X - Hex (lower/upper case)

	// Float formatting
	fmt.Printf("Float: %f\n", pi)       // %f - Decimal point, no exponent
	fmt.Printf("Scientific: %e\n", pi)  // %e - Scientific notation
	fmt.Printf("Precision: %.2f\n", pi) // %.2f - 2 decimal places

	// Boolean
	fmt.Printf("Boolean: %t\n", isStudent) // %t - true or false

	// Width and padding
	fmt.Printf("Padded: |%10s|%5d|\n", name, age) // Minimum width
	fmt.Printf("Left-aligned: |%-10s|\n", name)   // Left-aligned in width

	// Using Printf to format and print directly to stdout
	fmt.Printf("Name: %s, Age: %d\n", name, age) // Output: Name: Alice, Age: 30

	// Using Sprintf to create and return a formatted string
	// (useful when you need to store or further process the result)
	formattedString := fmt.Sprintf("Hello, my name is %s and I am %d years old.", name, age)
	println(formattedString) // Output: Hello, my name is Alice and I am 30 years old.

	// Other useful formatting functions:
	// - fmt.Fprintf: writes formatted output to a specified io.Writer
	// - fmt.Errorf: creates a formatted error
}
