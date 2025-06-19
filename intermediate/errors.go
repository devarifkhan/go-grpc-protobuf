package main

import "errors"

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("cannot compute square root of negative number")
	}
	return x * x, nil
}

func main() {
	result, err := sqrt(-4)
	if err != nil {
		println("Error:", err.Error())
	} else {
		println("Square root:", result)
	}

	result, err = sqrt(16)
	if err != nil {
		println("Error:", err.Error())
	} else {
		println("Square root:", result)
	}
}
