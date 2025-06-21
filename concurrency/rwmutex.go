package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var rwMutex sync.RWMutex
	var data = make(map[int]int)

	// Multiple readers can access the data concurrently
	for i := 0; i < 5; i++ {
		go func(id int) {
			for {
				rwMutex.RLock() // Acquire a read lock
				fmt.Printf("Reader %d: data = %v\n", id, data)
				rwMutex.RUnlock() // Release the read lock
				time.Sleep(time.Millisecond * 100)
			}
		}(i)
	}

	// Only one writer can access the data at a time
	// And no readers can access while a writer holds the lock
	for i := 0; i < 2; i++ {
		go func(id int) {
			counter := 0
			for {
				rwMutex.Lock() // Acquire a write lock
				counter++
				data[id] = counter
				fmt.Printf("Writer %d updated: data = %v\n", id, data)
				rwMutex.Unlock() // Release the write lock
				time.Sleep(time.Millisecond * 300)
			}
		}(i)
	}

	// Let the program run for a while
	time.Sleep(time.Second * 2)
}
