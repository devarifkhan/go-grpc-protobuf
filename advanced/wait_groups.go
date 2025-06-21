package main

// This file demonstrates the use of WaitGroups for goroutine synchronization.
// WaitGroups allow a program to wait for a collection of goroutines to finish
// executing before moving forward, which is essential for proper coordination
// in concurrent programs.

import (
	"fmt"  // For formatted output
	"sync" // For WaitGroup and other synchronization primitives
	"time" // For simulating work with time.Sleep
)

func main() {
	// Create a WaitGroup to synchronize goroutines
	// WaitGroup keeps track of how many goroutines are still running
	var wg sync.WaitGroup

	// Launch several goroutines and increment the WaitGroup counter for each
	for i := 1; i <= 5; i++ {
		// Add 1 to the WaitGroup counter BEFORE launching the goroutine
		// This is critical - adding after the goroutine starts can cause race conditions
		wg.Add(1)

		// Launch a worker goroutine with its ID and a pointer to the WaitGroup
		// Each worker will run concurrently with the main goroutine and other workers
		go waitGroupWorker(i, &wg)
	}

	// Block until the WaitGroup counter goes back to 0
	// This happens when all goroutines call wg.Done()
	// Without this wait, the program would exit before workers finish
	wg.Wait()

	fmt.Println("All workers completed their jobs")

	// At this point, we're guaranteed that all worker goroutines have finished
}

// waitGroupWorker simulates a worker that performs a task and signals completion
// Parameters:
//   - id: a unique identifier for the worker (for demonstration purposes)
//   - wg: a pointer to the WaitGroup for synchronization
func waitGroupWorker(id int, wg *sync.WaitGroup) {
	// Ensure we call Done when the function returns
	// defer guarantees this happens even if the function panics
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)

	// Simulate work with sleep
	// Each worker sleeps for a different duration (id seconds)
	time.Sleep(time.Second * time.Duration(id))

	fmt.Printf("Worker %d finished\n", id)

	// When this function returns, wg.Done() is automatically called (due to defer)
	// which decrements the WaitGroup counter
}
