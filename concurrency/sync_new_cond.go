package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	cond := sync.NewCond(&mutex)

	// Create a queue to simulate a shared resource
	queue := make([]int, 0, 10)

	// Producer goroutine
	go func() {
		for i := 0; i < 10; i++ {
			mutex.Lock()
			queue = append(queue, i)
			fmt.Printf("Produced: %d\n", i)
			cond.Signal() // Signal that a new item is available
			mutex.Unlock()
			time.Sleep(time.Millisecond * 100)
		}
	}()

	// Consumer goroutine
	go func() {
		for i := 0; i < 10; i++ {
			mutex.Lock()
			for len(queue) == 0 {
				fmt.Println("Queue empty, waiting...")
				cond.Wait() // Wait for signal that item is available
			}
			item := queue[0]
			queue = queue[1:]
			fmt.Printf("Consumed: %d\n", item)
			mutex.Unlock()
		}
	}()

	// Give time for goroutines to complete
	time.Sleep(time.Second * 2)
	fmt.Println("Program finished")
}
