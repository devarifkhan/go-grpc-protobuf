// Definition: A guessing game is a simple interactive program where the user tries to guess a value.
// How it works in Go: Uses input/output, random number generation, and control flow.
// Purpose: To demonstrate user input, random numbers, and control flow in Go.

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
