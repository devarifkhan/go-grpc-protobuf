package main

// This file implements the Leaky Bucket algorithm for rate limiting.
// The Leaky Bucket algorithm models a bucket with a small hole at the bottom
// that "leaks" at a constant rate. If the bucket is full, new requests are rejected.
// This creates a smoother flow of traffic by enforcing a consistent rate limit.

import (
	"fmt"
	"sync"
	"time"
)

// LeakyBucket implements a leaky bucket rate limiter
type LeakyBucket struct {
	capacity      int           // maximum number of tokens the bucket can hold
	leakRate      time.Duration // how frequently a token leaks (e.g. 100ms means 10 tokens per second)
	lastLeakTime  time.Time     // when the last token leaked
	currentTokens int           // current number of tokens in the bucket
	mutex         sync.Mutex    // for thread safety
}

// NewLeakyBucket creates a new leaky bucket rate limiter
// Parameters:
//   - capacity: Maximum number of tokens the bucket can hold
//   - leakRate: How frequently a token leaks out (e.g., 100ms means 10 tokens per second)
//
// Returns:
//   - A pointer to a new LeakyBucket instance
func NewLeakyBucket(capacity int, leakRate time.Duration) *LeakyBucket {
	return &LeakyBucket{
		capacity:      capacity,   // Maximum bucket size
		leakRate:      leakRate,   // How fast the bucket leaks
		lastLeakTime:  time.Now(), // Initialize leak time to now
		currentTokens: 0,          // Start with an empty bucket
	}
}

// Allow checks if a request is allowed based on the rate limit
// Unlike the typical leaky bucket implementation where tokens leak out at a constant rate,
// this implementation calculates the leakage that would have occurred since the last check
// Returns:
//   - true if the request is allowed
//   - false if the request should be throttled
func (lb *LeakyBucket) Allow() bool {
	// Lock for thread safety
	lb.mutex.Lock()
	defer lb.mutex.Unlock()

	// Calculate how many tokens should have leaked since last check
	now := time.Now()
	elapsed := now.Sub(lb.lastLeakTime)        // Time since last leak
	leakedTokens := int(elapsed / lb.leakRate) // How many tokens leaked out

	// Update token count and last leak time if any tokens leaked
	if leakedTokens > 0 {
		// Remove leaked tokens but don't go below zero
		lb.currentTokens = max(0, lb.currentTokens-leakedTokens)

		// Update the last leak time, preserving partial leaks
		// This accounts for the remainder of time that didn't result in a full token leak
		lb.lastLeakTime = now.Add(-elapsed % lb.leakRate)
	}

	// Check if we can add a new token (representing the current request)
	if lb.currentTokens < lb.capacity {
		// Bucket has room, add the token and allow the request
		lb.currentTokens++
		return true
	}

	// Bucket is full, reject the request
	return false
}

// max returns the larger of two integers
// This is a utility function used in our leaky bucket implementation
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Create a leaky bucket that allows 10 requests per second
	// - Capacity of 10 means it can handle bursts of up to 10 requests
	// - Leak rate of 100ms means a token leaks every 100ms (10 per second)
	limiter := NewLeakyBucket(10, 100*time.Millisecond)

	// Simulate 25 requests coming in at 50ms intervals (20 per second)
	// Since this exceeds our rate limit (10 per second), some will be rejected
	for i := 0; i < 25; i++ {
		// Try to add this request to the bucket
		allowed := limiter.Allow()
		fmt.Printf("Request %d: %v\n", i, allowed)

		// Small delay between requests (50ms = 20 requests per second)
		time.Sleep(50 * time.Millisecond)
	}

	// Expected output:
	// - First requests should be allowed until bucket fills up (10 requests)
	// - Then roughly every other request should be allowed as tokens leak
	//   (tokens leak at 100ms intervals, but requests arrive at 50ms intervals)
}
