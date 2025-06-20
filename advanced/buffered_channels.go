package main

import (
	"fmt"
	"time"
)

func main() {
	// Creating a buffered channel with capacity of 3
	ch := make(chan int, 3)

	// Sender goroutine
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("Sending: %d\n", i)
			ch <- i // This won't block until buffer is full
			fmt.Printf("Sent: %d\n", i)
		}
		close(ch) // Close the channel when done sending
	}()

	// Small delay to see the buffer effect
	time.Sleep(2 * time.Second)

	// Receive values from the channel
	for val := range ch {
		fmt.Printf("Received: %d\n", val)
		time.Sleep(1 * time.Second) // Simulate slow consumer
	}
}
