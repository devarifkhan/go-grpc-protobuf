package main

// This file demonstrates the use of buffered channels in Go.
// Buffered channels can store a limited number of values without a receiver ready,
// which provides asynchronous communication up to the capacity of the buffer.

import (
	"fmt"
	"time"
)

func main() {
	// Creating a buffered channel with capacity of 3
	// Unlike unbuffered channels, this can hold up to 3 values before blocking
	ch := make(chan int, 3)

	// Sender goroutine - demonstrates sending to a buffered channel
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("Sending: %d\n", i)
			// Send value to the channel
			// This won't block until the buffer is full (after 3 values)
			// Once the 4th value is sent, it will block until a value is received
			ch <- i
			fmt.Printf("Sent: %d\n", i)
		}
		// Close the channel when done sending to signal that no more values will be sent
		// This allows the range loop below to terminate when all values are received
		close(ch)
	}()

	// Small delay to see the buffer effect
	// This gives the sender time to fill the buffer before we start receiving
	time.Sleep(2 * time.Second)

	// Receive values from the channel using a range loop
	// This will continue until the channel is closed and all values are received
	for val := range ch {
		fmt.Printf("Received: %d\n", val)
		// Simulate a slow consumer to demonstrate that the sender can continue
		// sending until the buffer is full, even if the receiver is slow
		time.Sleep(1 * time.Second)
	}
}
