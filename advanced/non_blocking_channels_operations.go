package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// Non-blocking receive
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// Non-blocking send
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// Multi-way non-blocking select
	go func() {
		time.Sleep(time.Second)
		signals <- true
	}()

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

	// Using timeout in select
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case <-time.After(2 * time.Second):
		fmt.Println("timeout after 2 seconds")
	}

	fmt.Println("done")
}
