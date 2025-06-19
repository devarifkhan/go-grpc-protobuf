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