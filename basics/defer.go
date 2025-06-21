/*
DEFER IN GO
===========

Definition:
   The defer statement postpones the execution of a function until the surrounding function returns,
   either normally or through a panic.

How it works:
   1. Deferred functions are executed in LIFO (Last-In-First-Out) order
   2. Deferred functions run after the surrounding function completes but before it returns control
   3. Arguments to deferred functions are evaluated immediately when the defer statement is encountered

NOTE ON COMPILATION:
   Since this file is part of a larger project with multiple main functions,
   to run this specific example use:
   go run defer.go

   If you see a "main redeclared" error, you can:
   1. Run just this file: go run defer.go
   2. OR change the package name to something unique
   3. OR build it as part of a properly structured Go module

Common Use Cases:
   - Resource cleanup (closing files, network connections, database connections)
   - Unlocking mutexes
   - Ensuring functions are called even if errors occur

Benefits:
   - Keeps cleanup code close to resource acquisition code
   - Makes code more readable by keeping related operations together
   - Ensures cleanup happens even in error conditions
*/

package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// DeferExamples demonstrates various ways to use defer in Go
// To run this example, comment out the main function in other files or use:
// go run defer.go
func main() {
	fmt.Println("===== DEFER EXAMPLES =====")

	// Example 1: Basic defer usage with panic recovery
	fmt.Println("\nExample 1: Basic defer with panic recovery")
	processWithPanic()

	// Example 2: LIFO order demonstration
	fmt.Println("\nExample 2: LIFO order of deferred functions")
	deferOrder()

	// Example 3: Argument evaluation time
	fmt.Println("\nExample 3: Argument evaluation timing")
	deferArguments()

	// Example 4: Practical file handling example
	fmt.Println("\nExample 4: File handling with defer")
	fileHandlingExample()

	// Example 5: Using defer with mutexes
	fmt.Println("\nExample 5: Using defer with mutexes")
	mutexExample()

	// Example 6: Using defer for timing function execution
	fmt.Println("\nExample 6: Function execution timing")
	timeTrack()
}

// For demonstration purposes - uncomment this to run the examples
// func main() {
//     DeferExamples()
// }

// Example 1: Basic defer with panic recovery
func processWithPanic() {
	// This deferred function will execute even after a panic occurs
	defer fmt.Println("Deferred: This will always run last, even if an error occurs.")

	// This deferred function uses an anonymous function to recover from panic
	defer func() {
		// The recover() function stops the panic and returns the panic value
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Processing...")

	// Simulating an error with panic
	if true {
		panic("An error occurred!")
	}

	fmt.Println("This line will not be executed due to the panic.")
}

// Example 2: LIFO order demonstration
func deferOrder() {
	// Deferred functions are executed in Last-In-First-Out order
	defer fmt.Println("First defer (executes last)")
	defer fmt.Println("Second defer (executes second)")
	defer fmt.Println("Third defer (executes first)")

	fmt.Println("Function body (executes before any defers)")
}

// Example 3: Argument evaluation timing
func deferArguments() {
	i := 1

	// The argument '1' is evaluated immediately, but the function executes later
	defer fmt.Println("Value of i when defer encountered:", i)

	// Changing i doesn't affect the already evaluated defer argument
	i = 2
	fmt.Println("Value of i now:", i)
}

// Example 4: Practical file handling example
func fileHandlingExample() {
	// Create a temporary file for demonstration
	f, err := os.CreateTemp("", "example")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	// Defer file closing - happens automatically when function exits
	defer func() {
		fmt.Println("Closing file:", f.Name())
		f.Close()
		os.Remove(f.Name()) // Clean up by removing the temp file
	}()

	// Use the file
	fmt.Println("Working with file:", f.Name())
	fmt.Fprintln(f, "Writing some data to the file")

	// Even if we had an error here, the deferred function would still close the file
	// if err := someFunctionThatMightFail(); err != nil {
	//    return // The file would still be closed properly
	// }
}

// Example 5: Using defer with mutexes for thread safety
func mutexExample() {
	// Create a mutex to protect shared data
	var mu sync.Mutex
	counter := 0

	// Function that safely updates the counter
	increment := func() {
		// Lock the mutex
		mu.Lock()

		// Defer the unlock to ensure it happens even if the function panics
		defer mu.Unlock()

		// Critical section - access to shared data
		counter++
		fmt.Println("Counter incremented to:", counter)

		// Mutex will be automatically unlocked when function exits
		// This is cleaner than manually calling mu.Unlock() at every return point
	}

	// Execute the function a few times
	for i := 0; i < 3; i++ {
		increment()
	}

	fmt.Println("Final counter value:", counter)
	fmt.Println("Using defer with mutexes ensures locks are always released")
}

// Example 6: Using defer for timing function execution
func timeTrack() {
	// Create a function that measures execution time
	executionTimer := func(name string) func() {
		start := time.Now()
		fmt.Println("Starting operation:", name)

		// Return a function that prints the elapsed time
		return func() {
			elapsed := time.Since(start)
			fmt.Printf("Operation '%s' took %s to complete\n", name, elapsed)
		}
	}

	// Example of using the timer
	// The deferred function will execute at the end and report the time
	defer executionTimer("Slow operation")()

	// Simulate work
	fmt.Println("Performing slow operation...")
	time.Sleep(100 * time.Millisecond)

	// Nested timing example
	{
		defer executionTimer("Nested operation")()
		fmt.Println("Performing nested operation...")
		time.Sleep(50 * time.Millisecond)
	}

	fmt.Println("Using defer for timing provides clean start/end time measurement")
}
