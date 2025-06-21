package main

import (
	"fmt"
	"time"
)

func main() {
	// Create two channels
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Start goroutines that send messages to the channels
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Millisecond * 500)
			ch1 <- fmt.Sprintf("Message from channel 1: %d", i)
		}
		close(ch1)
	}()

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Millisecond * 800)
			ch2 <- fmt.Sprintf("Message from channel 2: %d", i)
		}
		close(ch2)
	}()

	// Use for-select to receive from both channels
	// until both are closed
	ch1Open, ch2Open := true, true
	for ch1Open || ch2Open {
		select {
		case msg, open := <-ch1:
			if open {
				fmt.Println(msg)
			} else {
				ch1Open = false
				fmt.Println("Channel 1 closed")
			}
		case msg, open := <-ch2:
			if open {
				fmt.Println(msg)
			} else {
				ch2Open = false
				fmt.Println("Channel 2 closed")
			}
		}
	}

	fmt.Println("All channels closed, exiting")
}
