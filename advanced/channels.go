package main

// This file demonstrates the basics of channels in Go.
// Channels are the pipes that connect concurrent goroutines, allowing
// them to communicate and synchronize their execution.
// This example shows both buffered and unbuffered channels.

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Create an unbuffered channel
	// Unbuffered channels require both sender and receiver to be ready at the same time
	// (synchronous communication)
	unbuffered := make(chan string)

	// Create a buffered channel with capacity 2
	// Buffered channels can hold up to the specified capacity without blocking
	// (asynchronous communication up to the buffer limit)
	buffered := make(chan string, 2)

	// WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// PART 1: Demonstrate unbuffered channel (requires sender and receiver to be ready)
	wg.Add(1)
	go func() {
		defer wg.Done() // Ensure the WaitGroup counter is decremented when done
		fmt.Println("Sender: Sending to unbuffered channel")
		// This will block until there's a receiver ready to receive the message
		// (demonstrating the synchronous nature of unbuffered channels)
		unbuffered <- "Hello from unbuffered"
		fmt.Println("Sender: Sent to unbuffered channel")
		// This line only executes after a receiver has taken the value
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second) // Simulate work before receiving
		fmt.Println("Receiver: About to receive from unbuffered")
		// This will unblock the sender when executed
		msg := <-unbuffered
		fmt.Println("Receiver:", msg)
	}()

	// PART 2: Demonstrate buffered channel
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Buffered: Sending first message")
		// These first two sends won't block because they fit in the buffer
		buffered <- "First"
		fmt.Println("Buffered: Sending second message")
		buffered <- "Second"
		// The buffer is now full (capacity 2)
		fmt.Println("Buffered: Sending third message (will block)")
		// This will block because the buffer is full and no one is receiving yet
		buffered <- "Third"
		// This line only executes after space is made in the buffer by a receiver
		fmt.Println("Buffered: Third message sent")
	}()

	// Wait a bit then read from buffered channel
	time.Sleep(3 * time.Second)
	fmt.Println("Main: Reading from buffered channel")
	// Receive all messages from the buffered channel
	// After the first receive, the blocked send of "Third" will complete
	fmt.Println(<-buffered) // "First"
	fmt.Println(<-buffered) // "Second"
	fmt.Println(<-buffered) // "Third"

	// Wait for all goroutines to complete before exiting
	wg.Wait()
}
