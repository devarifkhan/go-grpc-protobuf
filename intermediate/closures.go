package main

// This file demonstrates the concept of closures in Go.
// A closure is a function value that references variables from outside its body.
// The function may access and assign to the referenced variables; in this sense
// the function is "bound" to the variables.

func main() {
	// Example usage of the closure
	// Here we create a counter function that maintains its own state
	counter := createCounter()

	// Each call to counter() increments and returns the internal count
	// The count variable persists between calls because it's "closed over"
	println(counter()) // Output: 1
	println(counter()) // Output: 2
	println(counter()) // Output: 3

	// Creating a new counter would start from 0 again
	// This demonstrates that each closure has its own independent state
	counter2 := createCounter()
	println(counter2()) // Output: 1

	// The original counter continues from where it left off
	println(counter()) // Output: 4
}

// createCounter returns a function that increments and returns a counter value each time it is called.
// This demonstrates a practical use of closures to maintain state between function calls.
func createCounter() func() int {
	// This variable is declared outside the returned function
	// but is "captured" and persists between function calls
	count := 0 // Initial value

	// Return a closure (anonymous function) that can access the count variable
	// even after createCounter() has finished executing
	return func() int {
		count++      // Modify the captured variable
		return count // Return the new value

		// After this function returns, count still exists in memory because
		// the returned function maintains a reference to it
	}

	// Note: Each time createCounter is called, a new independent count variable
	// is created, so different counter functions will have separate states
}
