package main

// This file demonstrates the use of mutexes (mutual exclusion locks) in Go.
// Mutexes are used to protect shared resources from concurrent access,
// ensuring that only one goroutine can access the resource at a time.
// This prevents race conditions where multiple goroutines modify data simultaneously.

import (
	"fmt"
	"sync"
	"time"
)

// Counter is a simple counter with mutex protection
// This struct demonstrates a basic thread-safe data structure
type Counter struct {
	mu    sync.Mutex // Mutex to protect access to value
	value int        // The actual counter value
}

// Increment adds 1 to the counter in a thread-safe way
func (c *Counter) Increment() {
	// Lock the mutex before accessing the shared value
	// This ensures exclusive access - no other goroutines can access value while locked
	c.mu.Lock()

	// defer ensures the mutex is unlocked when the function returns
	// This happens even if the function panics, preventing deadlocks
	defer c.mu.Unlock()

	// Now it's safe to modify the value
	c.value++
}

// Value returns the current value of the counter in a thread-safe way
func (c *Counter) Value() int {
	// Lock the mutex for reading too, since reading while writing causes race conditions
	c.mu.Lock()
	defer c.mu.Unlock()

	// Return the value while mutex is still locked
	return c.value
	// The mutex will be unlocked after return due to defer
}

func main() {
	// Create a new counter instance
	counter := Counter{} // Initializes with value 0

	// Use WaitGroup to wait for all goroutines to complete
	var wg sync.WaitGroup

	// Launch 100 goroutines to increment counter concurrently
	// Without mutex protection, this would result in race conditions
	// and an unpredictable final value
	for i := 0; i < 100; i++ {
		// Add 1 to the WaitGroup counter before starting a goroutine
		wg.Add(1)

		// Launch a goroutine (anonymous function)
		go func() {
			// Ensure WaitGroup is decremented when goroutine completes
			defer wg.Done()

			// Call the thread-safe increment method
			counter.Increment()

			// Add some delay to simulate work and increase the chance
			// of concurrent access attempts (makes the example more realistic)
			time.Sleep(time.Millisecond)
		}()
	}

	// Wait for all goroutines to finish before continuing
	wg.Wait()

	// Print the final value - this should always be 100 because
	// we protected the counter with a mutex
	fmt.Println("Final counter value:", counter.Value())

	// Note: Without the mutex, the final value would likely be less than 100
	// due to race conditions where multiple goroutines read and increment
	// the value simultaneously
}
