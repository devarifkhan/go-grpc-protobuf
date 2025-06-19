package main

// interfaces defines the interface for the Person struct
type Person interface {
	Greet()
	HaveBirthday()
}

// PersonImpl implements the Person interface
type PersonImpl struct {
	Name string
	Age  int
}

// Greet prints a greeting message for the person.
func (p PersonImpl) Greet() {
	println("Hello, my name is", p.Name)
}

// HaveBirthday increments the person's age by 1 and prints a birthday message.
func (p *PersonImpl) HaveBirthday() {
	p.Age++
	println("Happy birthday! I am now", p.Age, "years old.")
}

func main() {
	// Create an instance of PersonImpl
	alice := PersonImpl{Name: "Alice", Age: 30}

	// Call methods on the struct
	alice.Greet()
	alice.HaveBirthday()

	// Create another person
	bob := PersonImpl{Name: "Bob", Age: 25}
	bob.Greet()
	bob.HaveBirthday()
	bob.Greet()                          // Name is unchanged
	println("Bob's new age is", bob.Age) // Age is updated because HaveBirthday uses a pointer receiver
}
