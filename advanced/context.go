package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create a context with cancellation capability
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Ensure resources are released

	// Create a context with timeout
	ctxTimeout, cancelTimeout := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelTimeout()

	// Create a context with deadline
	deadline := time.Now().Add(5 * time.Second)
	ctxDeadline, cancelDeadline := context.WithDeadline(context.Background(), deadline)
	defer cancelDeadline()

	// Create a context with value
	type key string
	ctxValue := context.WithValue(context.Background(), key("user"), "admin")

	// Example of using context in a goroutine
	go doWorkWithContext(ctx)

	// Example of handling timeout
	select {
	case <-ctxTimeout.Done():
		fmt.Println("Timeout occurred")
	case <-time.After(3 * time.Second):
		fmt.Println("This will not be printed because timeout occurs first")
	}

	// Get value from context
	if user, ok := ctxValue.Value(key("user")).(string); ok {
		fmt.Printf("User from context: %s\n", user)
	}

	// Give some time for the goroutine to run
	time.Sleep(1 * time.Second)
}

func doWorkWithContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work cancelled:", ctx.Err())
			return
		default:
			fmt.Println("Doing work...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
