package main

import (
	"fmt"
)

// This function only sends values on the channel (send-only)
func sendData(ch chan<- int) {
	fmt.Println("Sending data...")
	for i := 1; i <= 5; i++ {
		ch <- i
		fmt.Printf("Sent: %d\n", i)
	}
	close(ch)
}

// This function only receives values from the channel (receive-only)
func receiveData(ch <-chan int, done chan<- bool) {
	fmt.Println("Receiving data...")
	for value := range ch {
		fmt.Printf("Received: %d\n", value)
	}
	done <- true
}

func main() {
	// Create a bidirectional channel
	dataChannel := make(chan int)
	doneChannel := make(chan bool)

	// Start goroutines with specific channel directions
	go sendData(dataChannel)                 // dataChannel is passed as send-only
	go receiveData(dataChannel, doneChannel) // dataChannel is passed as receive-only

	// Wait for receiving to complete
	<-doneChannel
	fmt.Println("Program finished")
}
