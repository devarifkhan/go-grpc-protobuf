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
	http.HandleFunc("/teachers", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Teachers"))
		fmt.Println("Teachers Route")
	})

	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Students"))
		fmt.Println("Students Route")
	})

	http.HandleFunc("/exces", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Exces"))
		fmt.Println("Exces Route")
	})

	fmt.Print("Server listening on port ", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
