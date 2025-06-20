// Definition: Variables are named storage locations for data.
// How it works in Go: Declared using var or :=, with explicit or inferred types.
// Purpose: To demonstrate variable declaration, initialization, and usage in Go.

package main

import "fmt"

func main() {
	var age int = 30
	var name string = "John"
	var email string = "john.doe@example.com"

	fmt.Println("Age:", age)
	fmt.Println("Name:", name)
	fmt.Println("Email:", email)

	printName()
}

func printName() {
	firstName := "Doe"
	fmt.Println("Name:", firstName)
}
