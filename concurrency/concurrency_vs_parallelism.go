package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

func cpuIntensiveTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Task %d started\n", id)

	// Perform CPU-intensive work (calculating sine values in a loop)
	var result float64
	for i := 0; i < 10000000; i++ {
		result += math.Sin(float64(i))
	}

	fmt.Printf("Task %d completed (result: %.2f)\n", id, result)
}

func ioTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Task %d started\n", id)
	// Simulate I/O work with a sleep (non-CPU-intensive)
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("Task %d completed\n", id)
}

func main() {
	fmt.Println("Concurrency vs Parallelism in Go")

	// Check available CPUs
	cpus := runtime.NumCPU()
	fmt.Printf("Machine has %d CPU cores\n", cpus)

	// Default GOMAXPROCS value
	fmt.Printf("Default GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))

	// Demo with I/O bound tasks
	fmt.Println("\n=== Running I/O bound tasks ===")
	demoWithTasks("I/O-bound", ioTask)

	// Demo with CPU bound tasks
	fmt.Println("\n=== Running CPU bound tasks ===")
	demoWithTasks("CPU-bound", cpuIntensiveTask)
}

func demoWithTasks(taskType string, taskFunc func(int, *sync.WaitGroup)) {
	// Test with single CPU
	fmt.Printf("\n--- %s tasks with 1 CPU ---\n", taskType)
	runtime.GOMAXPROCS(1) // Force single-threaded execution
	runTasks(taskFunc)

	// Test with multiple CPUs
	fmt.Printf("\n--- %s tasks with %d CPUs ---\n", taskType, runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU()) // Use all available CPUs
	runTasks(taskFunc)
}

func runTasks(taskFunc func(int, *sync.WaitGroup)) {
	var wg sync.WaitGroup
	start := time.Now()

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go taskFunc(i, &wg)
	}

	wg.Wait()
	fmt.Printf("Execution took: %v\n", time.Since(start))
}
