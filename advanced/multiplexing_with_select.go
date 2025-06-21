package main

// This file demonstrates how to use the select statement in Go for channel multiplexing.
// Select allows a goroutine to wait on multiple communication operations.
// It blocks until one of its cases can run, then executes that case.
// If multiple cases are ready, one is chosen randomly.

import (
	"fmt"
	"time"
)

func main() {
	// Create two unbuffered channels for communication
	c1 := make(chan string)
	c2 := make(chan string)

	// Goroutine 1 - will send a message after 1 second
	go func() {
		// Simulate some work taking 1 second
		time.Sleep(1 * time.Second)
		// Send message to channel c1
		c1 <- "one"
	}()

	// Goroutine 2 - will send a message after 2 seconds
	go func() {
		// Simulate work taking 2 seconds
		time.Sleep(2 * time.Second)
		// Send message to channel c2
		c2 <- "two"
	}()

	// Use select to await both values simultaneously
	// We loop twice because we expect two messages (one from each channel)
	for i := 0; i < 2; i++ {
		// The select statement lets you wait on multiple channel operations
		// It will block until one of the cases can proceed, then it executes that case
		select {
		case msg1 := <-c1:
			// This case is selected when c1 has a message ready to receive
			fmt.Println("Received from channel 1:", msg1)
		case msg2 := <-c2:
			// This case is selected when c2 has a message ready to receive
			fmt.Println("Received from channel 2:", msg2)
		case <-time.After(3 * time.Second):
			// time.After creates a channel that sends a value after the specified duration
			// This is a timeout case - if neither c1 nor c2 send within 3 seconds
			fmt.Println("Timeout")
			return
		}
	}

	fmt.Println("All messages received")

	// Note: The order of messages received depends on the timing:
	// 1. Message from c1 (after 1 second)
	// 2. Message from c2 (after 2 seconds)
}
