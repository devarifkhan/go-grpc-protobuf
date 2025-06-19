package main

func main() {

	// Example usage of pointers
	a := 10
	b := &a // b is a pointer to a

	println("Value of a:", a)    // Output: Value of a: 10
	println("Address of a:", &a) // Output: Address of a: <some memory address>
	println("Value of b:", *b)   // Output: Value of b: 10

	*b = 20                       // Change the value at the address pointed by b
	println("New value of a:", a) // Output: New value of a: 20
}
