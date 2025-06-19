package main

func main() {

	// Example usage of structs
	type Person struct {
		Name string
		Age  int
	}

	// Create an instance of Person
	alice := Person{Name: "Alice", Age: 30}

	// Access fields of the struct
	println("Name:", alice.Name) // Output: Name: Alice
	println("Age:", alice.Age)   // Output: Age: 30

	// Modify fields of the struct
	alice.Age = 31
	println("Updated Age:", alice.Age) // Output: Updated Age: 31
}
