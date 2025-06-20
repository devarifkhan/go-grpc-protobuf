// Definition: switch is a control statement for multi-way branching based on the value of an expression.
// How it works in Go: Uses the switch keyword with cases and optional default.
// Purpose: To show how to use switch statements for cleaner branching logic.

package main

func main() {
	// This is a simple program that uses a switch statement
	// to print the day of the week based on a number.
	day := 3

	switch day {
	case 1:
		println("Monday")
	case 2:
		println("Tuesday")
	case 3:
		println("Wednesday")
	case 4:
		println("Thursday")
	case 5:
		println("Friday")
	case 6:
		println("Saturday")
	case 7:
		println("Sunday")
	default:
		println("Invalid day")
	}
}
