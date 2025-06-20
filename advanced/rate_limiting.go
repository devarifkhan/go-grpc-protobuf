package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// RateLimiter is a token bucket rate limiter
type RateLimiter struct {
	tokens      float64
	maxTokens   float64
	fillRate    float64
	mu          sync.Mutex
	lastRefresh time.Time
}

// NewRateLimiter creates a new rate limiter with specified max tokens and fill rate
func NewRateLimiter(maxTokens, fillRate float64) *RateLimiter {
	return &RateLimiter{
		tokens:      maxTokens,
		maxTokens:   maxTokens,
		fillRate:    fillRate,
		lastRefresh: time.Now(),
	}
}

// Allow returns true if the request can proceed
func (r *RateLimiter) Allow() bool {
	return r.AllowN(1)
}

// AllowN returns true if n requests can proceed
func (r *RateLimiter) AllowN(n float64) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(r.lastRefresh).Seconds()
	r.lastRefresh = now

	// Add tokens based on elapsed time
	r.tokens += elapsed * r.fillRate
	if r.tokens > r.maxTokens {
		r.tokens = r.maxTokens
	}

	// Check if enough tokens
	if r.tokens >= n {
		r.tokens -= n
		return true
	}
	return false
}

func main() {
	// Create rate limiter: 5 requests per second
	limiter := NewRateLimiter(5, 5)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	count := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Total allowed requests: %d in 2 seconds\n", count)
			return
		default:
			if limiter.Allow() {
				count++
				fmt.Println("Request allowed")
			} else {
				fmt.Println("Request denied")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
}
