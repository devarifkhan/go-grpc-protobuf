package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// An atomic counter accessible by all goroutines
	var counter uint64

	// We'll use a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Start 50 goroutines that each increment the counter 1000 times
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				// Atomically increment the counter
				atomic.AddUint64(&counter, 1)
			}
		}()
	}

	// Wait for all goroutines to complete
	wg.Wait()
	fmt.Println("Final Counter:", counter)

	// Example of other atomic operations
	var value uint64 = 10

	// Atomic load
	fmt.Println("Initial value:", atomic.LoadUint64(&value))

	// Atomic store
	atomic.StoreUint64(&value, 23)
	fmt.Println("After store:", atomic.LoadUint64(&value))

	// Compare and swap
	swapped := atomic.CompareAndSwapUint64(&value, 23, 100)
	fmt.Println("Swapped:", swapped, "Value now:", atomic.LoadUint64(&value))
}
