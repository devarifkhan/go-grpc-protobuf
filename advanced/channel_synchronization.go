package main

// This file demonstrates using channels for synchronization between goroutines.
// Channels are not just for communication but also for synchronization,
// allowing one goroutine to signal to another that processing is complete.

import (
	"fmt"
	"time"
)

// worker is a function that performs a task asynchronously
// and signals completion via a channel
func worker(done chan bool) {
	fmt.Print("Working...")

	// Simulate work by sleeping for a second
	// In a real application, this would be actual processing
	time.Sleep(time.Second)

	fmt.Println("Done")

	// Send notification on channel when done
	// This signal will be received by the main goroutine
	done <- true
}

func main() {
	// Create a channel to synchronize execution
	// The buffer size of 1 means we can send one value
	// without the sender blocking
	done := make(chan bool, 1)

	// Start worker goroutine to run concurrently with main
	// worker will perform its task asynchronously
	go worker(done)

	// Block until worker sends a notification on done channel
	// This is the synchronization point - main will wait here until
	// the worker goroutine signals that it's finished
	<-done

	fmt.Println("Main function finished")
}
