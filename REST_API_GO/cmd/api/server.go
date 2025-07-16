package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type user struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	City string `json:"city"`
}

func main() {
	port := ":3000"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "Hello World")
		//w.Write([]byte("Hello World"))
		fmt.Println("Hello World Route")
	})
	http.HandleFunc("/teachers", func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(r.Method)
		switch r.Method {
		case http.MethodGet:
			_, err := w.Write([]byte("Hello GET Method on Teachers Route"))
			if err != nil {
				return
			}
			fmt.Println("Hello GET Method on Teachers Route")
		case http.MethodPost:
			//parse the request body (x-www-form-urlencoded)
			//err := r.ParseForm()
			//if err != nil {
			//	http.Error(w, "Error parsing form", http.StatusBadRequest)
			//	fmt.Println("Error parsing form:", err)
			//	return
			//}
			//fmt.Println("form", r.Form)
			//
			////prepare the response data
			//response := make(map[string]interface{})
			//for key, value := range r.Form {
			//	response[key] = value[0]
			//}

			// RAW Body (json)
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

			// if you expect JSON data, you can unmarshal it
			var userInstance user
			err = json.Unmarshal(body, &userInstance)
			if err != nil {
				return
			}
			fmt.Println("Unmarshalled JSON into a user struct:", userInstance)
			fmt.Println("User Name:", userInstance.Name)
			fmt.Println("User Age:", userInstance.Age)
			fmt.Println("User City:", userInstance.City)

			// prepare the response
			response := make(map[string]interface{})
			for key, value := range r.Form {
				response[key] = value[0]
			}

			err = json.Unmarshal(body, &response)
			if err != nil {
				http.Error(w, "Error unmarshalling body", http.StatusBadRequest)
				fmt.Println("Error unmarshalling body:", err)
				return
			}

			fmt.Println("Unmarshalled JSON into a map:", response)

			//w.Write([]byte("Hello POST Method on Teachers Route"))
			fmt.Println("Hello POST Method on Teachers Route")
		case http.MethodPut:
			//w.Write([]byte("Hello PUT Method on Teachers Route"))
			fmt.Println("Hello PUT Method on Teachers Route")
		case http.MethodPatch:
			//w.Write([]byte("Hello PATCH Method on Teachers Route"))
			fmt.Println("Hello PATCH Method on Teachers Route")
		case http.MethodDelete:
			//w.Write([]byte("Hello DELETE Method on Teachers Route"))
			fmt.Println("Hello DELETE Method on Teachers Route")

		}

		_, err := w.Write([]byte("Teachers"))
		if err != nil {
			return
		}
		fmt.Println("Teachers Route")
	})

	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Students"))
		if err != nil {
			return
		}
		fmt.Println("Students Route")
	})

	http.HandleFunc("/exces", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Exces"))
		if err != nil {
			return
		}
		fmt.Println("Exces Route")
	})

	fmt.Print("Server listening on port ", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
