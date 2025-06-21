package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	var wg sync.WaitGroup

	// Function that will be executed only once
	onceFunc := func() {
		fmt.Println("This will be printed only once")
	}

	// Launch multiple goroutines
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d trying to execute the once function\n", id)
			once.Do(onceFunc) // This will only execute the first time it's called
			fmt.Printf("Goroutine %d finished\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("All goroutines completed")
}
