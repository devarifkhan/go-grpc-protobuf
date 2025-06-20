package main

// hashing_crypto
import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
)

func main() {
	// Read the file content
	data, err := ioutil.ReadFile("bufio.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create a new SHA256 hash
	hash := sha256.New()

	// Write data to the hash
	hash.Write(data)

	// Get the final hash result
	hashInBytes := hash.Sum(nil)

	// Convert the byte slice to a hexadecimal string
	hashString := hex.EncodeToString(hashInBytes)

	fmt.Println("SHA256 Hash:", hashString)
}
