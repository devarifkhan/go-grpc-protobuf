// Definition: Constants are immutable values which are known at compile time and do not change during execution.
// How it works in Go: Declared using the const keyword, constants can be typed or untyped.
// Purpose: To illustrate the declaration and use of constants in Go.

package main

const pi = 3.14
const e = 2.71828

func main() {
	// This is a simple program that uses a constant
	// to print the value of pi.
	println("The value of pi is:", pi)
	println("The value of e is:", e)

}
