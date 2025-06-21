package main

// This file demonstrates how to properly close channels in Go.
// Closing channels is important to signal to receivers that no more data will be sent.
// It also demonstrates how receivers can detect when a channel is closed.

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Create an unbuffered channel
	ch := make(chan int)

	// WaitGroup to ensure both goroutines finish before the program exits
	var wg sync.WaitGroup
	wg.Add(2) // Add 2 to the counter since we have two goroutines

	// Sender goroutine - produces values and eventually closes the channel
	go func() {
		// Ensure the WaitGroup counter is decremented when the goroutine exits
		defer wg.Done()

		// Ensure the channel is closed when the goroutine exits
		// Important: Only the sender should close a channel, never the receiver
		// Closing a channel signals that no more values will be sent
		defer close(ch)

		// Send 5 values through the channel
		for i := 1; i <= 5; i++ {
			ch <- i
			time.Sleep(100 * time.Millisecond) // Simulate work between sends
		}
		fmt.Println("Sender completed, closing channel")
		// The channel will be closed here due to the deferred close(ch)
	}()

	// Receiver goroutine - consumes values until the channel is closed
	go func() {
		defer wg.Done()

		// Infinite loop that breaks when the channel is closed
		for {
			// The second return value (ok) indicates whether the channel is still open
			val, ok := <-ch
			if !ok {
				// ok == false means the channel is closed and empty
				fmt.Println("Channel closed, receiver exiting")
				return
			}
			fmt.Printf("Received: %d\n", val)
		}

		// Alternative way to receive until the channel is closed:
		// for val := range ch {
		//     fmt.Printf("Received: %d\n", val)
		// }
		// No need to check if channel is closed with range
	}()

	// Wait for both goroutines to finish before exiting
	wg.Wait()
	fmt.Println("Program completed")
}
