package main

import "fmt"

func main(){
	var age int = 30
	var name string = "John"
	var email string = "john.doe@example.com"

	fmt.Println("Age:", age)
	fmt.Println("Name:", name)
	fmt.Println("Email:", email)

	printName()
}

func printName() {
	firstName := "Doe"
	fmt.Println("Name:", firstName)
}

