package main

import (
	"fmt"
	"time"
)

// worker is a function that does some work and then notifies when it's done
func worker(done chan bool) {
	fmt.Print("Working...")
	time.Sleep(time.Second) // Simulate work
	fmt.Println("Done")

	// Send notification on channel when done
	done <- true
}

func main() {
	// Create a channel to synchronize execution
	done := make(chan bool, 1)

	// Start worker goroutine
	go worker(done)

	// Block until worker sends a notification on done channel
	<-done
	fmt.Println("Main function finished")
}
