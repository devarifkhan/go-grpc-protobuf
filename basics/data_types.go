// Definition: Data types specify the type of data a variable can hold, such as int, float, string, etc.
// How it works in Go: Go is statically typed, so each variable must have a type, either explicitly or inferred.
// Purpose: To explain the basic data types available in Go.

package main

import "fmt"

func main() {
	fmt.Println("Hello" + "World!")
	fmt.Println("9 x 10 = ", 9*10)
	fmt.Println("180.18/2.0 = ", 100.18/2.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
