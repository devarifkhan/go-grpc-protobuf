// Definition: Arithmetic operators perform mathematical operations like addition, subtraction, multiplication, division, and modulus.
// How it works in Go: Go supports standard arithmetic operators (+, -, *, /, %) for numeric types.
// Purpose: To demonstrate how to use arithmetic operators in Go.

package main

import "fmt"

func main() {
	// This is a simple program that uses arithmetic operators
	// to perform basic calculations.
	a := 10
	b := 5

	sum := a + b
	diff := a - b
	prod := a * b
	quot := a / b
	mod := a % b

	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", diff)
	fmt.Println("Product:", prod)
	fmt.Println("Quotient:", quot)
	fmt.Println("Modulus:", mod)
}
