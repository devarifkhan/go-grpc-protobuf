package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Create a channel
	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	// Sender goroutine
	go func() {
		defer wg.Done()
		defer close(ch) // Sender is responsible for closing the channel

		for i := 1; i <= 5; i++ {
			ch <- i
			time.Sleep(100 * time.Millisecond)
		}
		fmt.Println("Sender completed, closing channel")
	}()

	// Receiver goroutine
	go func() {
		defer wg.Done()

		for {
			val, ok := <-ch
			if !ok {
				// Channel is closed
				fmt.Println("Channel closed, receiver exiting")
				return
			}
			fmt.Printf("Received: %d\n", val)
		}
	}()

	wg.Wait()
	fmt.Println("Program completed")
}
