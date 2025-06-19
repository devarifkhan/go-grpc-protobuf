package main

// example for panic in go
func main() {
	defer func() {
		if r := recover(); r != nil {
			println("Recovered from panic:", r)
		}
	}()

	panic("This is a panic message")
}
