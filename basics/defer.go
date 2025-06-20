// Definition: The defer statement postpones the execution of a function until the surrounding function returns.
// How it works in Go: Deferred functions are executed in LIFO order after the main function completes.
// Purpose: To demonstrate resource cleanup and function call ordering using defer.

package main

import "fmt"

func main() {
	process()

}

func process() {
	defer fmt.Println("Deferred: This will always run last, even if an error occurs.")

	fmt.Println("Processing...")

	// Simulating an error
	if true {
		panic("An error occurred!")
	}

	fmt.Println("This line will not be executed due to the panic.")
}
