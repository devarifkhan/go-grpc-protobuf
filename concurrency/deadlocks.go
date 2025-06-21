package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Create two mutex locks
	var lock1, lock2 sync.Mutex

	// Create a wait group to wait for goroutines to finish
	var wg sync.WaitGroup
	wg.Add(2)

	// First goroutine: locks lock1, then tries to lock lock2
	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 1: Acquiring lock1...")
		lock1.Lock()
		fmt.Println("Goroutine 1: Acquired lock1")

		// Sleep to increase the chance of deadlock
		time.Sleep(200 * time.Millisecond)

		fmt.Println("Goroutine 1: Attempting to acquire lock2...")
		lock2.Lock()
		fmt.Println("Goroutine 1: Acquired lock2") // This will never execute

		// If we get here, release the locks
		lock2.Unlock()
		lock1.Unlock()
	}()

	// Second goroutine: locks lock2, then tries to lock lock1
	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 2: Acquiring lock2...")
		lock2.Lock()
		fmt.Println("Goroutine 2: Acquired lock2")

		// Sleep to increase the chance of deadlock
		time.Sleep(200 * time.Millisecond)

		fmt.Println("Goroutine 2: Attempting to acquire lock1...")
		lock1.Lock()
		fmt.Println("Goroutine 2: Acquired lock1") // This will never execute

		// If we get here, release the locks
		lock1.Unlock()
		lock2.Unlock()
	}()

	// Wait for both goroutines (this will wait forever due to deadlock)
	fmt.Println("Main: Waiting for goroutines...")
	wg.Wait()
	fmt.Println("Main: Done") // This will never execute
}
