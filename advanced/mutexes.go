package main

import (
	"fmt"
	"sync"
	"time"
)

// Counter is a simple counter with mutex protection
type Counter struct {
	mu    sync.Mutex
	value int
}

// Increment adds 1 to the counter in a thread-safe way
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value returns the current value of the counter
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	// Create a new counter
	counter := Counter{}

	// Use WaitGroup to wait for all goroutines to complete
	var wg sync.WaitGroup

	// Launch 100 goroutines to increment counter
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
			time.Sleep(time.Millisecond) // Add some delay to simulate work
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Print the final value
	fmt.Println("Final counter value:", counter.Value())
}
