// Definition: A for loop is a control structure for repeated execution of a block of code.
// How it works in Go: Go uses the for keyword for all loop types (traditional, while-like, infinite).
// Purpose: To show different ways to use for loops in Go.

package main

func main() {
	// This is a simple program that uses a for loop
	// to print the numbers from 1 to 10.
	for i := 1; i <= 10; i++ {
		println("Number:", i)
	}

	// This is a simple program that uses a for loop
	// to print the numbers from 10 to 1.
	for i := 10; i >= 1; i-- {
		println("Number:", i)
	}

	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		println("Index:", index, "Value:", value)
	}
}
