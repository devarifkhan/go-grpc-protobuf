package main

// This file implements the Fixed Window Counter algorithm for rate limiting.
// The Fixed Window Counter algorithm divides time into fixed windows
// (e.g., 1-second intervals) and allows a maximum number of requests in each window.
// While simple to implement, this approach can allow twice the rate limit
// if requests are clustered at window boundaries.

import (
	"context" // For controlled cancellation and timeouts
	"fmt"     // For formatted output
	"sync"    // For mutual exclusion (thread safety)
	"time"    // For time-related functions
)

// FixedWindowRateLimiter implements a simple fixed-window rate limiting strategy
// This approach tracks request counts within fixed time windows and resets
// the counter when a new time window begins.
type FixedWindowRateLimiter struct {
	mu            sync.Mutex    // For thread safety across goroutines
	windowSize    time.Duration // Size of each time window (e.g., 1 second)
	maxRequests   int           // Maximum number of requests allowed per window
	currentWindow time.Time     // Start time of the current window
	counter       int           // Number of requests in the current window
}

// NewFixedWindowRateLimiter creates a new rate limiter with given window size and max requests
// Parameters:
//   - windowSize: The duration of each time window, e.g., 1 second
//   - maxRequests: The maximum number of requests allowed in each window
//
// Returns:
//   - A pointer to a new FixedWindowRateLimiter instance
func NewFixedWindowRateLimiter(windowSize time.Duration, maxRequests int) *FixedWindowRateLimiter {
	return &FixedWindowRateLimiter{
		windowSize:    windowSize,  // Set the window size (e.g., 1 second)
		maxRequests:   maxRequests, // Set maximum allowed requests per window
		currentWindow: time.Now(),  // Initialize the first window to start now
		// counter starts at 0 by default
	}
}

// Allow checks if a request should be allowed based on the rate limit
// This method is thread-safe and can be called from multiple goroutines
// Returns:
//   - true if the request is allowed to proceed
//   - false if the request exceeds the rate limit and should be rejected
func (rl *FixedWindowRateLimiter) Allow() bool {
	// Lock to ensure thread safety when checking/updating counters
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()

	// If we're in a new time window, reset the counter
	// This happens when current time minus window start time exceeds the window size
	if now.Sub(rl.currentWindow) >= rl.windowSize {
		// Reset counter for the new window
		rl.counter = 0
		// Update the current window start time
		rl.currentWindow = now
	}

	// If counter is less than max requests, increment and allow the request
	if rl.counter < rl.maxRequests {
		rl.counter++
		return true // Request is allowed
	}

	// Request exceeds the limit for the current window
	return false // Request is rejected

	// Note: This implementation has an edge case at window boundaries
	// Requests could bunch up at the end of one window and beginning of another,
	// potentially allowing up to 2x the desired rate briefly
}

func main() {
	// Create a rate limiter with 5 requests per second window
	limiter := NewFixedWindowRateLimiter(1*time.Second, 5)

	// PART 1: Simulate 10 concurrent requests (bursting)
	// This demonstrates how the limiter handles concurrent requests in the same window
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done() // Ensure WaitGroup counter is decremented when goroutine exits

			// Try to get permission from the rate limiter
			allowed := limiter.Allow()
			fmt.Printf("Request %d allowed: %v\n", id, allowed)

			// Expected result: First 5 requests (in any order) will be allowed
			// and the remaining 5 will be rejected since they exceed the window limit
		}(i)
	}
	// Wait for all concurrent requests to complete
	wg.Wait()

	// PART 2: Simulate requests spread over time
	// This demonstrates how the rate limiter works across multiple time windows
	fmt.Println("\nSimulating requests over time:")

	// Create a context that will automatically cancel after 3 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // Ensure resources are released when function exits

	// Create a ticker that "ticks" every 200ms (5 times per second)
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop() // Ensure ticker is stopped when function exits

	count := 0
	for {
		select {
		case <-ctx.Done():
			// Context is cancelled (timeout or explicit cancellation)
			return
		case <-ticker.C:
			// A tick has occurred (every 200ms)
			count++

			// Try to get permission from the rate limiter
			allowed := limiter.Allow()
			fmt.Printf("Time: %v, Request %d allowed: %v\n",
				time.Now().Format("15:04:05.000"), count, allowed)

			// Expected pattern: Since we have a 1-second window with max 5 requests,
			// and we're sending requests every 200ms (5 per second),
			// we should see the first 5 requests in each second allowed and
			// the rest rejected until the window resets
		}
	}
}
