package main

import (
	"fmt"
	"testing"
)

// Function to benchmark
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// This is a main function to demonstrate the functionality
func main() {
	fmt.Println("Fibonacci(10):", Fibonacci(10))

	// Note: Benchmarks are not run from main
	// They should be executed using 'go test -bench=.'
}

// BenchmarkFibonacci benchmarks the Fibonacci function
// This would go in a separate *_test.go file in practice
func BenchmarkFibonacci(b *testing.B) {
	// Run the Fibonacci function b.N times
	for i := 0; i < b.N; i++ {
		Fibonacci(10)
	}
}

// An alternative implementation for benchmarking comparison
func FibonacciMemoized(n int) int {
	memo := make(map[int]int)
	var fib func(int) int

	fib = func(n int) int {
		if n <= 1 {
			return n
		}

		if val, ok := memo[n]; ok {
			return val
		}

		memo[n] = fib(n-1) + fib(n-2)
		return memo[n]
	}

	return fib(n)
}

// Benchmark for the memoized version
func BenchmarkFibonacciMemoized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciMemoized(10)
	}
}
