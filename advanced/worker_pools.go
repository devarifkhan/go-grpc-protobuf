package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID          int
	Description string
}

type Result struct {
	JobID  int
	Output string
}

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job.ID)
		time.Sleep(time.Second) // Simulate work
		output := fmt.Sprintf("Processed job %d: %s", job.ID, job.Description)
		results <- Result{JobID: job.ID, Output: output}
		fmt.Printf("Worker %d finished job %d\n", id, job.ID)
	}
}

func main() {
	numJobs := 10
	numWorkers := 3

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	// Start workers
	var wg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j, Description: fmt.Sprintf("Task #%d", j)}
	}
	close(jobs)

	// Wait for workers to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	for result := range results {
		fmt.Printf("Result: %s\n", result.Output)
	}
}
