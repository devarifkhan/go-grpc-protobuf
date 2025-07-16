package main

import (
	"fmt"
	"io"
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

		fmt.Println(r.Method)
		switch r.Method {
		case http.MethodGet:
			w.Write([]byte("Hello GET Method on Teachers Route"))
			fmt.Println("Hello GET Method on Teachers Route")
		case http.MethodPost:
			//parse the request body
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Error parsing form", http.StatusBadRequest)
				fmt.Println("Error parsing form:", err)
				return
			}
			fmt.Println("form", r.Form)

			//prepare the response data
			response := make(map[string]interface{})
			for key, value := range r.Form {
				response[key] = value[0]
			}

			// RAW Body
			body, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Error reading body", http.StatusBadRequest)
				fmt.Println("Error reading body:", err)
				return
			}
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {

				}
			}(r.Body)
			fmt.Println("Raw Body:", string(body))

			fmt.Println("Processed Response Map:", response)
			w.Write([]byte("Hello POST Method on Teachers Route"))
			fmt.Println("Hello POST Method on Teachers Route")
		case http.MethodPut:
			w.Write([]byte("Hello PUT Method on Teachers Route"))
			fmt.Println("Hello PUT Method on Teachers Route")
		case http.MethodPatch:
			w.Write([]byte("Hello PATCH Method on Teachers Route"))
			fmt.Println("Hello PATCH Method on Teachers Route")
		case http.MethodDelete:
			w.Write([]byte("Hello DELETE Method on Teachers Route"))
			fmt.Println("Hello DELETE Method on Teachers Route")

		}

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
