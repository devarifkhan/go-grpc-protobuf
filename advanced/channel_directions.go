package main

// This file demonstrates channel direction constraints in Go.
// Go allows specifying a channel as send-only or receive-only,
// which provides better type safety and makes the intention clearer.

import (
	"fmt"
)

// sendData only sends values on the channel (send-only)
// The chan<- syntax indicates this is a send-only channel from the function's perspective
// This ensures the function cannot read from the channel, preventing potential bugs
func sendData(ch chan<- int) {
	fmt.Println("Sending data...")
	for i := 1; i <= 5; i++ {
		// Send value to the channel
		ch <- i
		fmt.Printf("Sent: %d\n", i)
	}
	// Close the channel to signal that no more values will be sent
	// Note: Only the sender should close a channel, never the receiver
	close(ch)
}

// receiveData only receives values from the channel (receive-only)
// The <-chan syntax indicates this is a receive-only channel from the function's perspective
// This ensures the function cannot write to or close the channel
func receiveData(ch <-chan int, done chan<- bool) {
	fmt.Println("Receiving data...")
	// Range over the channel to receive all values until it's closed
	for value := range ch {
		fmt.Printf("Received: %d\n", value)
	}
	// Signal that all values have been processed by sending true on the done channel
	done <- true
}

func main() {
	// Create a bidirectional channel that can be both read from and written to
	// In the main function, we have full access to the channel in both directions
	dataChannel := make(chan int)

	// Create another channel to signal when processing is complete
	doneChannel := make(chan bool)

	// Start goroutines with specific channel directions
	// Although dataChannel is bidirectional in main, we pass it with direction constraints
	// to the goroutines to enforce how they can use it
	go sendData(dataChannel)                 // dataChannel is passed as send-only (chan<-)
	go receiveData(dataChannel, doneChannel) // dataChannel is passed as receive-only (<-chan)

	// Wait for receiving to complete
	// This blocks main from exiting until a value is received on doneChannel
	<-doneChannel
	fmt.Println("Program finished")
}
