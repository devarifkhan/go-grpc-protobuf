package main

// This file demonstrates goroutines - lightweight threads managed by the Go runtime.
// Goroutines allow concurrent execution of functions and are a fundamental
// building block for concurrent programming in Go.

import (
	"fmt"
	"sync"
	"time"
)

// worker simulates a worker that performs a task
// id: a unique identifier for the worker
// wg: a pointer to WaitGroup to signal when the worker completes
func worker(id int, wg *sync.WaitGroup) {
	// Ensure the WaitGroup counter is decremented when the function returns
	// defer ensures this happens even if the function panics
	defer wg.Done()

	// Log when the worker starts
	fmt.Printf("Worker %d starting\n", id)

	// Simulate work being done
	time.Sleep(time.Second)

	// Log when the worker completes its work
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// WaitGroup is used to wait for a collection of goroutines to finish
	var wg sync.WaitGroup

	// Launch several goroutines
	for i := 1; i <= 5; i++ {
		// Increment the WaitGroup counter before starting each goroutine
		// This tells the WaitGroup how many goroutines to wait for
		wg.Add(1)

		// Start the worker in a new goroutine
		// The 'go' keyword launches a new goroutine that runs concurrently
		// with the calling goroutine (in this case, main)
		go worker(i, &wg)
	}

	// Wait for all goroutines to complete
	// This blocks until the WaitGroup counter becomes zero
	// (when all workers have called wg.Done())
	wg.Wait()

	fmt.Println("All workers completed")

	// At this point, all goroutines have finished execution
	// The program will exit only when the main goroutine finishes
}
