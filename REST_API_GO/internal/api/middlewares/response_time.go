package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func ResponseTimeMiddleware(next http.Handler) http.Handler {
	fmt.Println("Response Time Middleware...")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("---Response Time Middleware start---")
		start := time.Now()
		wrappedWritter := &responseWritter{ResponseWriter: w, status: http.StatusOK}

		// calculate the response time
		duration := time.Since(start)
		w.Header().Set("X-Response-Time", duration.String())
		next.ServeHTTP(wrappedWritter, r)

		// log the request details
		duration = time.Since(start)
		log.Printf("Method %s, URL %s, Status %d, Duration %s", r.Method, r.URL.Path, wrappedWritter.status, duration)
		fmt.Println("---Response Time Middleware end---")

	})
}

type responseWritter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWritter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}
