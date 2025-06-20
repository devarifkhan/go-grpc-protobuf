package main

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
func NewLeakyBucket(capacity int, leakRate time.Duration) *LeakyBucket {
	return &LeakyBucket{
		capacity:      capacity,
		leakRate:      leakRate,
		lastLeakTime:  time.Now(),
		currentTokens: 0,
	}
}

// Allow checks if a request is allowed based on the rate limit
func (lb *LeakyBucket) Allow() bool {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()

	// Calculate how many tokens should have leaked since last check
	now := time.Now()
	elapsed := now.Sub(lb.lastLeakTime)
	leakedTokens := int(elapsed / lb.leakRate)

	// Update token count and last leak time
	if leakedTokens > 0 {
		lb.currentTokens = max(0, lb.currentTokens-leakedTokens)
		lb.lastLeakTime = now.Add(-elapsed % lb.leakRate)
	}

	// Check if we can add a new token
	if lb.currentTokens < lb.capacity {
		lb.currentTokens++
		return true
	}

	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Create a leaky bucket that allows 10 requests per second
	limiter := NewLeakyBucket(10, 100*time.Millisecond)

	// Simulate requests
	for i := 0; i < 25; i++ {
		allowed := limiter.Allow()
		fmt.Printf("Request %d: %v\n", i, allowed)

		// Small delay between requests
		time.Sleep(50 * time.Millisecond)
	}
}
