package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name    string
	Age     int
	Address string
}

func main() {
	// Create a Person instance
	p := Person{
		Name:    "Alice",
		Age:     30,
		Address: "123 Main St",
	}

	// Get the reflect.Value of p
	v := reflect.ValueOf(p)

	// Get the reflect.Type of p
	t := reflect.TypeOf(p)

	fmt.Println("Type:", t.Name())
	fmt.Println("Kind:", t.Kind())

	// Iterate through the struct fields
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		fmt.Printf("%s: %v (type: %v)\n", field.Name, value, field.Type)
	}

	// Create a new value using reflection
	newP := reflect.New(t).Elem()
	newP.Field(0).SetString("Bob")
	newP.Field(1).SetInt(25)
	newP.Field(2).SetString("456 Oak Ave")

	fmt.Println("\nNew person:", newP.Interface().(Person))
}
