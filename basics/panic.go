// Definition: panic is used to stop the normal execution of a Go program when an unexpected error occurs.
// How it works in Go: Calling panic causes the program to crash and run deferred functions.
// Purpose: To demonstrate error handling and program termination in Go.

package main

// example for panic in go
func main() {
	defer func() {
		if r := recover(); r != nil {
			println("Recovered from panic:", r)
		}
	}()

	panic("This is a panic message")
}
