package main

// This file demonstrates the use of the context package in Go.
// Context is used to carry deadlines, cancellation signals, and other request-scoped values
// across API boundaries and between processes. It's particularly useful for controlling
// the lifetime of operations spanning multiple goroutines.

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 1. Context with cancellation capability
	// WithCancel returns a copy of parent with a new Done channel
	// The returned context's Done channel is closed when the returned cancel function is called
	ctx, cancel := context.WithCancel(context.Background())
	// Always call cancel when you're done to release resources
	defer cancel()

	// 2. Context with timeout
	// WithTimeout returns a copy of parent but with a timeout duration
	// The context will be cancelled automatically after the timeout expires
	ctxTimeout, cancelTimeout := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelTimeout() // Good practice to defer cancel even with timeout
	// 3. Context with deadline
	// WithDeadline returns a copy of parent but with an absolute expiration time
	// Different from timeout which is a relative duration
	expirationTime := time.Now().Add(5 * time.Second)
	ctxDeadline, cancelDeadline := context.WithDeadline(context.Background(), expirationTime)
	defer cancelDeadline()

	// Use ctxDeadline to prevent linting error
	deadlineTime, _ := ctxDeadline.Deadline()
	fmt.Println("Deadline set for:", deadlineTime)

	// 4. Context with value
	// WithValue returns a copy of parent in which the value associated with key is val
	// Used to pass request-scoped values across API boundaries
	type key string
	ctxValue := context.WithValue(context.Background(), key("user"), "admin")

	// Example of using context to control a goroutine
	// The goroutine will monitor the context's Done channel
	go doWorkWithContext(ctx)

	// Example of handling timeout using select
	// This demonstrates how to respond when a context expires
	select {
	case <-ctxTimeout.Done():
		fmt.Println("Timeout occurred")
	case <-time.After(3 * time.Second):
		fmt.Println("This will not be printed because timeout occurs first")
	}

	// Retrieving a value from context
	// Type assertion (.(string)) is required because Value returns interface{}
	if user, ok := ctxValue.Value(key("user")).(string); ok {
		fmt.Printf("User from context: %s\n", user)
	}

	// Give some time for the goroutine to run before program exits
	time.Sleep(1 * time.Second)

	// At this point, the cancel() deferred above will be called,
	// signaling to any functions using ctx that they should stop
}

// doWorkWithContext demonstrates how to make a function respond to context cancellation
// This pattern is common in servers, long-running operations, or any code that needs
// to be responsive to cancellation requests
func doWorkWithContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			// The Done channel is closed when the context is cancelled or times out
			// ctx.Err() returns the reason why the context was cancelled
			fmt.Println("Work cancelled:", ctx.Err()) // Will print "context canceled"
			return
		default:
			// Continue working until the context is cancelled
			fmt.Println("Doing work...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
