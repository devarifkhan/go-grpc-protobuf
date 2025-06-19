package main

// Person represents a person with a name and age
type Person struct {
	Name string
	Age  int
}

// methods
func main() {
	// Create an instance of Person
	alice := Person{Name: "Alice", Age: 30}

	// Call methods on the struct
	alice.Greet()
	alice.HaveBirthday()
}

// Greet prints a greeting message for the person.
func (p Person) Greet() {
	println("Hello, my name is", p.Name)
}

// HaveBirthday increments the person's age by 1 and prints a birthday message.
func (p *Person) HaveBirthday() {
	p.Age++
	println("Happy birthday! I am now", p.Age, "years old.")
}
