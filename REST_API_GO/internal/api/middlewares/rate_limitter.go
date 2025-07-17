package middlewares

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
)

type rateLimitter struct {
	mu        sync.Mutex
	visitors  map[string]int
	limit     int
	resetTime time.Duration
}

func NewRateLimitter(limit int, resetTime time.Duration) *rateLimitter {
	r1 := &rateLimitter{
		visitors:  make(map[string]int),
		limit:     limit,
		resetTime: resetTime,
	}
	go r1.resetVisitorCounts()
	return r1
}

func (r *rateLimitter) resetVisitorCounts() {
	for {
		time.Sleep(r.resetTime)
		r.mu.Lock()
		for k := range r.visitors {
			delete(r.visitors, k)
		}
		r.mu.Unlock()
	}
}

func (r *rateLimitter) Middleware(next http.Handler) http.Handler {
	fmt.Println("HPP Middleware...")
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("---HPP Middleware start---")
		r.mu.Lock()
		defer r.mu.Unlock()

		ip, _, err := net.SplitHostPort(req.RemoteAddr)
		if err != nil {
			ip = req.RemoteAddr // fallback if parsing fails
		}

		if _, exists := r.visitors[ip]; !exists {
			r.visitors[ip] = 0
		}

		if r.visitors[ip] >= r.limit {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		r.visitors[ip]++
		next.ServeHTTP(w, req)
		fmt.Println("---HPP Middleware end---")
	})
}
