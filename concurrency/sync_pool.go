package main

import (
	"fmt"
	"sync"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// Create a new sync.Pool with a New function that creates Person objects
	personPool := sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating a new Person")
			return &Person{}
		},
	}

	// Get a Person from the pool
	person1 := personPool.Get().(*Person)
	person1.Name = "Alice"
	person1.Age = 25
	fmt.Printf("Person 1: %+v\n", person1)

	// Return the Person to the pool when done
	personPool.Put(person1)

	// Get another Person from the pool (reuses the one we Put back)
	person2 := personPool.Get().(*Person)
	fmt.Printf("Person 2 before reset: %+v\n", person2)

	// Reset fields for reuse
	person2.Name = "Bob"
	person2.Age = 30
	fmt.Printf("Person 2 after reset: %+v\n", person2)

	// Demonstrate parallel usage with multiple goroutines
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Get from pool
			p := personPool.Get().(*Person)
			p.Name = fmt.Sprintf("Person %d", id)
			p.Age = 20 + id

			fmt.Printf("Goroutine %d: %+v\n", id, p)

			// Return to pool
			personPool.Put(p)
		}(i)
	}

	wg.Wait()
}
