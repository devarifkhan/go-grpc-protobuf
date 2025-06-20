package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Create a channel to receive signals
	sigs := make(chan os.Signal, 1)

	// Register for specific signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Create a channel to indicate when we're done
	done := make(chan bool, 1)

	// Start a goroutine to handle signals
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println("Received signal:", sig)
		done <- true
	}()

	fmt.Println("Program is running. Press Ctrl+C to terminate.")
	fmt.Println("PID:", os.Getpid())

	// Wait for signal
	<-done
	fmt.Println("Gracefully shutting down")
}
