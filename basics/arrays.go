// Definition: An array is a fixed-size, ordered collection of elements of the same type.
// How it works in Go: Arrays are declared with a fixed length and type, e.g., var arr [5]int.
// Purpose: To show how to declare, initialize, and use arrays in Go.

package main

func main() {
	// This is a simple program that uses a while loop
	// to print the numbers from 1 to 10.
	i := 1
	for i <= 10 {
		println("Number:", i)
		i++
	}

	// This is a simple program that uses a while loop
	// to print the numbers from 10 to 1.
	j := 10
	for j >= 1 {
		println("Number:", j)
		j--
	}

	numbers := []int{1, 2, 3, 4, 5}
	index := 0
	for index < len(numbers) {
		println("Index:", index, "Value:", numbers[index])
		index++
	}
}
