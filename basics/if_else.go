// Definition: if-else is a conditional statement for branching logic based on boolean expressions.
// How it works in Go: Uses if, else if, and else keywords for decision making.
// Purpose: To demonstrate conditional execution in Go.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 1 and 100
	target := rand.Intn(100) + 1
	var guess int

	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("I have selected a number between 1 and 100.")

	// Loop until the user guesses the correct number
	for {
		fmt.Print("Enter your guess: ")
		fmt.Scan(&guess)

		if guess < target {
			fmt.Println("Too low! Try again.")
		} else if guess > target {
			fmt.Println("Too high! Try again.")
		} else {
			fmt.Println("Congratulations! You've guessed the number:", target)
			break
		}
	}
}
