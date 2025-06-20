package main

import (
	"fmt"
	"sort"
)

func main() {
	// Sorting integers
	numbers := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(numbers)
	fmt.Println("Sorted integers:", numbers)

	// Sorting strings
	names := []string{"charlie", "alice", "bob", "david"}
	sort.Strings(names)
	fmt.Println("Sorted strings:", names)

	// Sorting floats
	scores := []float64{3.5, 1.2, 5.8, 2.3}
	sort.Float64s(scores)
	fmt.Println("Sorted floats:", scores)

	// Custom sorting with sort.Slice
	people := []struct {
		Name string
		Age  int
	}{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 22},
	}

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("Sorted by age:", people)

	// Sort by name
	sort.Slice(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})
	fmt.Println("Sorted by name:", people)
}
