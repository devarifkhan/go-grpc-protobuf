package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Example of a race condition
	counter := 0

	// Create a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Launch 1000 goroutines that increment the counter
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			// Each goroutine reads counter, increments it, and writes it back
			temp := counter
			temp++
			time.Sleep(time.Nanosecond) // Force context switching to make race condition more likely
			counter = temp
			wg.Done()
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("Counter with race condition:", counter)

	// Fix with mutex
	counterSafe := 0
	var mutex sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			mutex.Lock()
			counterSafe++
			mutex.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter with mutex protection:", counterSafe)
}
