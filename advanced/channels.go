package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Create an unbuffered channel
	unbuffered := make(chan string)

	// Create a buffered channel with capacity 2
	buffered := make(chan string, 2)

	var wg sync.WaitGroup

	// Demonstrate unbuffered channel (requires sender and receiver to be ready)
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Sender: Sending to unbuffered channel")
		unbuffered <- "Hello from unbuffered"
		fmt.Println("Sender: Sent to unbuffered channel")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second) // Simulate work
		fmt.Println("Receiver: About to receive from unbuffered")
		msg := <-unbuffered
		fmt.Println("Receiver:", msg)
	}()

	// Demonstrate buffered channel
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Buffered: Sending first message")
		buffered <- "First"
		fmt.Println("Buffered: Sending second message")
		buffered <- "Second"
		fmt.Println("Buffered: Sending third message (will block)")
		buffered <- "Third" // This will block until someone reads
		fmt.Println("Buffered: Third message sent")
	}()

	// Wait a bit then read from buffered channel
	time.Sleep(3 * time.Second)
	fmt.Println("Main: Reading from buffered channel")
	fmt.Println(<-buffered)
	fmt.Println(<-buffered)
	fmt.Println(<-buffered)

	wg.Wait()
}
