package main

// This file demonstrates non-blocking operations on channels using the select statement.
// By default, channel operations (send and receive) are blocking, meaning they will wait
// until the other side is ready. Non-blocking operations allow checking channels without
// waiting, implementing timeouts, and handling multiple channels simultaneously.

import (
	"fmt"
	"time"
)

func main() {
	// Create two channels for our examples
	messages := make(chan string) // Channel for string messages
	signals := make(chan bool)    // Channel for boolean signals

	// EXAMPLE 1: Non-blocking receive operation
	// This will immediately check if a message is available without waiting
	select {
	case msg := <-messages:
		// This case executes if a message is available to receive
		fmt.Println("Received message:", msg)
	default:
		// This default case executes immediately if no message is available
		// Without the default, this would block until a message arrives
		fmt.Println("No message received")
	}

	// EXAMPLE 2: Non-blocking send operation
	// This will try to send a message without blocking if the receiver isn't ready
	msg := "hi"
	select {
	case messages <- msg:
		// This case executes if the message can be sent (someone is ready to receive)
		fmt.Println("Sent message:", msg)
	default:
		// This default case executes immediately if the send cannot proceed
		// Without the default, this would block until a receiver is available
		fmt.Println("No message sent")
	}

	// EXAMPLE 3: Multi-way non-blocking select
	// Start a goroutine that will send a signal after 1 second
	go func() {
		time.Sleep(time.Second) // Wait for 1 second
		signals <- true         // Send a signal
	}()

	// Try to receive from either channel, or proceed immediately if neither is ready
	select {
	case msg := <-messages:
		// This case executes if a message is available on the messages channel
		fmt.Println("Received message:", msg)
	case sig := <-signals:
		// This case executes if a signal is available on the signals channel
		fmt.Println("Received signal:", sig)
	default:
		// This executes if neither channel has data ready
		fmt.Println("No activity")
	}

	// EXAMPLE 4: Using timeout in select
	// This will wait for a message, but only up to a specified timeout
	select {
	case msg := <-messages:
		// This case executes if a message arrives before the timeout
		fmt.Println("Received message:", msg)
	case <-time.After(2 * time.Second):
		// time.After returns a channel that sends a value after the specified duration
		// This case executes if no message is received within 2 seconds
		fmt.Println("Timeout after 2 seconds")
	}
	// Note: No default case here, so this select will block until either
	// a message is received or the timeout occurs

	fmt.Println("Done")
}
