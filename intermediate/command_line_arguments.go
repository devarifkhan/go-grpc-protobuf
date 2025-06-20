package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args provides access to raw command-line arguments
	// os.Args[0] is the program name
	// os.Args[1:] are the arguments

	fmt.Println("Program name:", os.Args[0])

	if len(os.Args) > 1 {
		fmt.Println("Arguments:")
		for i, arg := range os.Args[1:] {
			fmt.Printf("  %d: %s\n", i+1, arg)
		}
	} else {
		fmt.Println("No command line arguments provided")
	}
}
