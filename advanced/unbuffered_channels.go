package main

import (
	"fmt"
	"time"
)

func main() {
	// Create an unbuffered channel
	ch := make(chan int)

	// Start a goroutine that sends data to the channel
	go func() {
		fmt.Println("Goroutine: About to send value")
		ch <- 42 // This will block until someone receives
		fmt.Println("Goroutine: Value sent")
	}()

	// Wait a bit to demonstrate blocking behavior
	time.Sleep(2 * time.Second)

	fmt.Println("Main: About to receive value")
	value := <-ch // This unblocks the sender
	fmt.Println("Main: Received value:", value)

	// Give time for the goroutine to finish printing
	time.Sleep(time.Millisecond * 100)

	fmt.Println("Program complete")
}
