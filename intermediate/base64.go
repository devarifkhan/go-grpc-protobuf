package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("bufio.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println(encoded)
}
