package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v int64
}

// Inc increments the counter by one.
func (c *SafeCounter) Inc() {
	atomic.AddInt64(&c.v, 1)
}

// Value returns the current value of the counter.
func (c *SafeCounter) Value() int64 {
	return atomic.LoadInt64(&c.v)
}

func main() {
	// Create a stateful counter
	counter := SafeCounter{}

	// Simulate work in multiple goroutines
	for i := 0; i < 100; i++ {
		go func() {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			counter.Inc()
		}()
	}

	// Wait for all goroutines to finish
	time.Sleep(2 * time.Second)

	fmt.Println("Final counter value:", counter.Value())
}
