package main

//read write files
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("bufio.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(strings.ToUpper(line))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
