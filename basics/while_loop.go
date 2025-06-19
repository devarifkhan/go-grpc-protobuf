package main

func main() {

	num :=1
	for num <= 10 {
		println("Number:", num)
		num++
	}
	// This is a simple program that uses a while loop
	// to print the numbers from 10 to 1.
	for num := 10; num >= 1; num-- {
		println("Number:", num)
	}
	// This is a simple program that uses a while loop
	// to print the numbers from 1 to 5.
	numbers := []int{1, 2, 3, 4, 5}
	for index := 0; index < len(numbers); index++ {
		println("Index:", index, "Value:", numbers[index])
	}
}