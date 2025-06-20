package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// FixedWindowRateLimiter implements a simple fixed-window rate limiting strategy
type FixedWindowRateLimiter struct {
	mu            sync.Mutex
	windowSize    time.Duration
	maxRequests   int
	currentWindow time.Time
	counter       int
}

// NewFixedWindowRateLimiter creates a new rate limiter with given window size and max requests
func NewFixedWindowRateLimiter(windowSize time.Duration, maxRequests int) *FixedWindowRateLimiter {
	return &FixedWindowRateLimiter{
		windowSize:    windowSize,
		maxRequests:   maxRequests,
		currentWindow: time.Now(),
	}
}

// Allow checks if a request should be allowed based on the rate limit
func (rl *FixedWindowRateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	// If we're in a new time window, reset the counter
	if now.Sub(rl.currentWindow) >= rl.windowSize {
		rl.counter = 0
		rl.currentWindow = now
	}

	// If counter is less than max requests, increment and allow
	if rl.counter < rl.maxRequests {
		rl.counter++
		return true
	}

	return false
}

func main() {
	// Create a rate limiter with 5 requests per second
	limiter := NewFixedWindowRateLimiter(1*time.Second, 5)

	// Simulate 10 concurrent requests
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			allowed := limiter.Allow()
			fmt.Printf("Request %d allowed: %v\n", id, allowed)
		}(i)
	}
	wg.Wait()

	// Simulate requests over time
	fmt.Println("\nSimulating requests over time:")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	count := 0
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			count++
			allowed := limiter.Allow()
			fmt.Printf("Time: %v, Request %d allowed: %v\n", time.Now().Format("15:04:05.000"), count, allowed)
		}
	}
}
