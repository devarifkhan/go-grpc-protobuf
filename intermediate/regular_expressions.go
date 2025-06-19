package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Example string
	str := "Hello, World!"

	// Match a simple pattern
	match, _ := regexp.MatchString("Hello", str)
	fmt.Println("Match found:", match)

	// Find all occurrences of a pattern
	re := regexp.MustCompile("[aeiou]")
	matches := re.FindAllString(str, -1)
	fmt.Println("Vowels found:", matches)

	// Replace all occurrences of a pattern
	replaced := re.ReplaceAllString(str, "*")
	fmt.Println("Replaced string:", replaced)
}
