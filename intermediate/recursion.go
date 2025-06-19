package main

func main() {

	// Example usage of the factorial function
	result := factorial(5)
	println(result) // Output: 120
}

// factorial calculates the factorial of a given number n using recursion.
func factorial(n int) int {
	if n == 0 {
		return 1 // Base case: factorial of 0 is 1
	}
	return n * factorial(n-1) // Recursive case: n * factorial of (n-1)
}
