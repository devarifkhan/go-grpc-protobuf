package main

import (
	"fmt"
	"sync"
	"time"
)

// TokenBucket represents a token bucket rate limiter
type TokenBucket struct {
	capacity     int
	tokens       int
	refillRate   int // tokens per second
	lastRefillTS time.Time
	mutex        sync.Mutex
}

// NewTokenBucket creates a new token bucket rate limiter
func NewTokenBucket(capacity, refillRate int) *TokenBucket {
	return &TokenBucket{
		capacity:     capacity,
		tokens:       capacity,
		refillRate:   refillRate,
		lastRefillTS: time.Now(),
		mutex:        sync.Mutex{},
	}
}

// Allow checks if a request can proceed
func (tb *TokenBucket) Allow() bool {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()

	tb.refill()
	if tb.tokens > 0 {
		tb.tokens--
		return true
	}
	return false
}

// refill adds tokens based on elapsed time
func (tb *TokenBucket) refill() {
	now := time.Now()
	elapsed := now.Sub(tb.lastRefillTS)
	tokensToAdd := int(elapsed.Seconds()) * tb.refillRate

	if tokensToAdd > 0 {
		tb.tokens = min(tb.capacity, tb.tokens+tokensToAdd)
		tb.lastRefillTS = now
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	limiter := NewTokenBucket(5, 2) // 5 tokens capacity, refills 2 tokens per second

	for i := 0; i < 10; i++ {
		if limiter.Allow() {
			fmt.Println("Request allowed:", i)
		} else {
			fmt.Println("Request throttled:", i)
		}
		time.Sleep(250 * time.Millisecond)
	}
}
