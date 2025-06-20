package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// Launch several goroutines and increment the WaitGroup counter for each
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// Block until the WaitGroup counter goes back to 0
	// This happens when all goroutines call wg.Done()
	wg.Wait()
	fmt.Println("All workers completed their jobs")
}

func worker(id int, wg *sync.WaitGroup) {
	// Ensure we call Done when the function returns
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)
	// Simulate work with sleep
	time.Sleep(time.Second * time.Duration(id))
	fmt.Printf("Worker %d finished\n", id)
}
