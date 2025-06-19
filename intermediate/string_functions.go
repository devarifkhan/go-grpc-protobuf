package main

// StringFunctions defines a set of string manipulation functions
type StringFunctions struct{}

// Reverse returns the reversed version of the input string
func (sf StringFunctions) Reverse(s string) string {
	reversed := ""
	for i := len(s) - 1; i >= 0; i-- {
		reversed += string(s[i])
	}
	return reversed
}

// ToUpper converts the input string to uppercase
func (sf StringFunctions) ToUpper(s string) string {
	upper := ""
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			upper += string(char - 32) // Convert to uppercase
		} else {
			upper += string(char)
		}
	}
	return upper
}

// ToLower converts the input string to lowercase
func (sf StringFunctions) ToLower(s string) string {
	lower := ""
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			lower += string(char + 32) // Convert to lowercase
		} else {
			lower += string(char)
		}
	}
	return lower
}

// Contains checks if the input string contains the specified substring
func (sf StringFunctions) Contains(s, substr string) bool {
	return len(s) >= len(substr) && s[:len(substr)] == substr
}

// Count returns the number of occurrences of the specified substring in the input string
func (sf StringFunctions) Count(s, substr string) int {
	count := 0
	substrLen := len(substr)
	for i := 0; i <= len(s)-substrLen; i++ {
		if s[i:i+substrLen] == substr {
			count++
		}
	}
	return count
}

// Trim removes leading and trailing whitespace from the input string
func (sf StringFunctions) Trim(s string) string {
	start := 0
	end := len(s) - 1

	// Find the first non-whitespace character
	for start <= end && (s[start] == ' ' || s[start] == '\t' || s[start] == '\n') {
		start++
	}

	// Find the last non-whitespace character
	for end >= start && (s[end] == ' ' || s[end] == '\t' || s[end] == '\n') {
		end--
	}

	if start > end {
		return ""
	}
	return s[start : end+1]
}

// Split splits the input string by the specified delimiter
func (sf StringFunctions) Split(s, delimiter string) []string {
	if delimiter == "" {
		return []string{s} // Return the original string if delimiter is empty
	}

	var result []string
	start := 0

	for i := 0; i <= len(s)-len(delimiter); i++ {
		if s[i:i+len(delimiter)] == delimiter {
			result = append(result, s[start:i])
			start = i + len(delimiter)
		}
	}

	// Add the last segment after the last delimiter
	if start < len(s) {
		result = append(result, s[start:])
	}

	return result
}

// Join joins the elements of the input slice into a single string with the specified delimiter
func (sf StringFunctions) Join(elements []string, delimiter string) string {
	if len(elements) == 0 {
		return ""
	}

	result := elements[0]
	for _, elem := range elements[1:] {
		result += delimiter + elem
	}
	return result
}

// Replace replaces all occurrences of the old substring with the new substring in the input string
func (sf StringFunctions) Replace(s, old, new string) string {
	if old == "" {
		return s // If old is empty, return the original string
	}

	result := ""
	start := 0
	oldLen := len(old)

	for i := 0; i <= len(s)-oldLen; i++ {
		if s[i:i+oldLen] == old {
			result += s[start:i] + new
			start = i + oldLen
		}
	}

	// Add the remaining part of the string after the last occurrence
	result += s[start:]

	return result
}

func main() {
	sf := StringFunctions{}

	// Example string
	str := "Hello, World!"

	// Reverse the string
	reversed := sf.Reverse(str)
	println("Reversed:", reversed)

	// Convert to uppercase
	upper := sf.ToUpper(str)
	println("Uppercase:", upper)

	// Convert to lowercase
	lower := sf.ToLower(str)
	println("Lowercase:", lower)

	// Check if the string contains a substring
	contains := sf.Contains(str, "World")
	println("Contains 'World':", contains)

	// Count occurrences of a substring
	count := sf.Count(str, "o")
	println("Count of 'o':", count)

	// Trim whitespace
	trimmed := sf.Trim("   Hello, World!   ")
	println("Trimmed:", trimmed)

	// Split the string
	split := sf.Split(str, ", ")
	println("Split:", split[0], "|", split[1])

	// Join elements into a string
	joined := sf.Join([]string{"Hello", "World"}, ", ")
	println("Joined:", joined)

	// Replace substring
	replaced := sf.Replace(str, "World", "Go")
	println("Replaced:", replaced)
}
