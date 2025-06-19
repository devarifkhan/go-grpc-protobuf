package main

func swap[T any](a, b *T) {
	// Swap the values of a and b
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	a := 5
	b := 10
	swap(&a, &b)
	println("After swap:", a, b)
}
