package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := ":3000"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "Hello World")
		w.Write([]byte("Hello World"))
		fmt.Println("Hello World Route")
	})

	fmt.Print("Server listening on port ", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
