package main

// This file implements the Token Bucket algorithm for rate limiting.
// The Token Bucket algorithm is a common rate limiting strategy that allows
// bursts of traffic up to a configurable limit while maintaining a steady
// long-term rate. It works by adding tokens to a bucket at a fixed rate and
// allowing requests to proceed only if a token is available.

import (
	"fmt"
	"sync"
	"time"
)

// TokenBucket represents a token bucket rate limiter
// It limits the rate of operations by requiring tokens from a conceptual "bucket"
// that is filled at a constant rate. Operations can proceed if there are enough tokens.
type TokenBucket struct {
	capacity     int        // Maximum number of tokens the bucket can hold
	tokens       int        // Current number of tokens in the bucket
	refillRate   int        // Rate at which tokens are added (tokens per second)
	lastRefillTS time.Time  // Timestamp of the last token refill
	mutex        sync.Mutex // Mutex to ensure thread safety during token operations
}

// NewTokenBucket creates a new token bucket rate limiter
// Parameters:
//   - capacity: The maximum number of tokens the bucket can hold
//   - refillRate: How many tokens to add per second
//
// Returns:
//   - A pointer to a new TokenBucket instance, initialized with full capacity
func NewTokenBucket(capacity, refillRate int) *TokenBucket {
	return &TokenBucket{
		capacity:     capacity,     // Set maximum capacity
		tokens:       capacity,     // Start with a full bucket of tokens
		refillRate:   refillRate,   // Set token refill rate
		lastRefillTS: time.Now(),   // Initialize refill timestamp
		mutex:        sync.Mutex{}, // Initialize mutex for thread safety
	}
}

// Allow checks if a request can proceed by attempting to consume one token
// Returns:
//   - true if a token was available and consumed (request can proceed)
//   - false if no tokens are available (request should be throttled)
//
// This method is thread-safe and can be called concurrently from multiple goroutines
func (tb *TokenBucket) Allow() bool {
	// Lock the mutex to ensure atomic operations on the token bucket
	tb.mutex.Lock()
	// Ensure the mutex is unlocked when the function returns
	defer tb.mutex.Unlock()

	// First refill any tokens based on elapsed time
	tb.refill()

	// Check if we have tokens available
	if tb.tokens > 0 {
		// Consume one token
		tb.tokens--
		return true // Request is allowed
	}
	return false // No tokens available, request is throttled
}

// refill adds tokens to the bucket based on elapsed time since the last refill
// This is an internal method that should only be called when the mutex is locked
func (tb *TokenBucket) refill() {
	// Get current time
	now := time.Now()

	// Calculate time elapsed since last refill
	elapsed := now.Sub(tb.lastRefillTS)

	// Calculate how many tokens to add based on elapsed time and refill rate
	// The cast to int truncates any fractional seconds
	tokensToAdd := int(elapsed.Seconds()) * tb.refillRate

	// Only update if we're adding tokens
	if tokensToAdd > 0 {
		// Add tokens but don't exceed capacity
		tb.tokens = min(tb.capacity, tb.tokens+tokensToAdd)

		// Update timestamp of last refill
		tb.lastRefillTS = now
	}
}

// min returns the smaller of two integers
// This is a utility function to ensure we don't exceed bucket capacity
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Create a new token bucket limiter:
	// - Capacity of 5 tokens (allows bursts of up to 5 requests)
	// - Refill rate of 2 tokens per second (long-term rate limit)
	limiter := NewTokenBucket(5, 2)

	// Simulate 10 requests coming in at a rate of 4 per second (every 250ms)
	// This exceeds our refill rate of 2 per second, so some should be throttled
	for i := 0; i < 10; i++ {
		// Try to get a token for this request
		if limiter.Allow() {
			fmt.Println("Request allowed:", i) // Token obtained, request proceeds
		} else {
			fmt.Println("Request throttled:", i) // No token available, request throttled
		}

		// Wait before next request (4 requests per second)
		time.Sleep(250 * time.Millisecond)
	}

	// Expected behavior:
	// - First 5 requests allowed (initial bucket capacity)
	// - Then some throttling as tokens refill at 2 per second
}
