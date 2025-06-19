package main

func main() {

	// Example usage of the closure
	counter := createCounter()
	println(counter()) // Output: 1
	println(counter()) // Output: 2
	println(counter()) // Output: 3
}

// createCounter returns a function that increments and returns a counter value each time it is called.
func createCounter() func() int {
	count := 0 // This variable is captured by the closure

	// Return a function that increments and returns the count
	return func() int {
		count++
		return count
	}
}
