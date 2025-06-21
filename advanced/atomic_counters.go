// Package main demonstrates the use of atomic operations in Go.
// Atomic operations are used for synchronization between goroutines
// without using mutexes, making them more efficient for simple operations.
package main

// This file demonstrates the use of atomic operations in Go.
// Atomic operations are used for synchronization between goroutines
// without using mutexes, making them more efficient for simple operations.

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// An atomic counter accessible by all goroutines
	// Atomic operations allow safe concurrent access without locks
	var counter uint64

	// We'll use a WaitGroup to wait for all goroutines to finish
	// WaitGroup is used to wait for a collection of goroutines to finish execution
	var wg sync.WaitGroup

	// Start 50 goroutines that each increment the counter 1000 times
	// This demonstrates concurrent access to a shared counter using atomic operations
	for i := 0; i < 50; i++ {
		// Add 1 to the WaitGroup counter before starting a goroutine
		wg.Add(1)

		// Launch a goroutine (concurrent function)
		go func() {
			// Ensure the WaitGroup counter is decremented when the goroutine exits
			defer wg.Done()

			// Each goroutine performs 1000 increments
			for j := 0; j < 1000; j++ {
				// Atomically increment the counter
				// AddUint64 atomically adds delta to *addr and returns the new value
				// This ensures no race conditions occur between goroutines
				atomic.AddUint64(&counter, 1)
			}
		}()
	}

	// Wait for all goroutines to complete
	// This blocks until the WaitGroup counter becomes zero
	wg.Wait()
	fmt.Println("Final Counter:", counter) // Should be exactly 50,000

	// Example of other atomic operations
	var value uint64 = 10

	// Atomic load operation
	// LoadUint64 atomically loads *addr without any risk of another goroutine modifying it during the read
	fmt.Println("Initial value:", atomic.LoadUint64(&value))

	// Atomic store operation
	// StoreUint64 atomically stores val into *addr without risk of race conditions
	atomic.StoreUint64(&value, 23)
	fmt.Println("After store:", atomic.LoadUint64(&value))

	// Compare and swap operation
	// CompareAndSwapUint64 executes the compare-and-swap operation for an uint64 value
	// It only updates the value if the current value matches the expected value
	swapped := atomic.CompareAndSwapUint64(&value, 23, 100)
	fmt.Println("Swapped:", swapped, "Value now:", atomic.LoadUint64(&value))
}
