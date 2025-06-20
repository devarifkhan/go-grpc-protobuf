package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	// Goroutine 1
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	// Goroutine 2
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// Use select to await both values simultaneously
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received from channel 1:", msg1)
		case msg2 := <-c2:
			fmt.Println("Received from channel 2:", msg2)
		case <-time.After(3 * time.Second):
			fmt.Println("Timeout")
			return
		}
	}

	fmt.Println("All messages received")
}
