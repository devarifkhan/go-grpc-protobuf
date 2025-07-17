package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	mw "restapi/internal/api/middlewares"
)

type user struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	City string `json:"city"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello World"))
	if err != nil {
		return
	}
	fmt.Println("Hello World Route")
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Method)
	switch r.Method {
	case http.MethodGet:
		_, err := w.Write([]byte("Hello GET Method on Teachers Route"))
		if err != nil {
			return
		}

		fmt.Println("Hello GET Method on Teachers Route")
	case http.MethodPost:
		_, err := w.Write([]byte("Hello POST Method on Teachers Route"))
		if err != nil {
			return
		}
		fmt.Println("Hello POST Method on Teachers Route")
	case http.MethodPut:
		_, err := w.Write([]byte("Hello PUT Method on Teachers Route"))
		if err != nil {
			return
		}
		fmt.Println("Hello PUT Method on Teachers Route")
	case http.MethodPatch:
		_, err := w.Write([]byte("Hello PATCH Method on Teachers Route"))
		if err != nil {
			return
		}
		fmt.Println("Hello PATCH Method on Teachers Route")
	case http.MethodDelete:
		_, err := w.Write([]byte("Hello DELETE Method on Teachers Route"))
		if err != nil {
			return
		}
		fmt.Println("Hello DELETE Method on Teachers Route")

	}

	_, err := w.Write([]byte("Teachers"))
	if err != nil {
		return
	}

}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Students"))
	if err != nil {
		return
	}
	fmt.Println("Students Route")
}

func ExcesHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Exces"))
	if err != nil {
		return
	}
	fmt.Println("Exces Route")
}
func main() {
	port := ":3000"

	cert := "cert.pem"
	key := "key.pem"

	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/teachers/", teachersHandler)

	mux.HandleFunc("/students/", studentsHandler)

	http.HandleFunc("/exces/", ExcesHandler)

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	server := &http.Server{
		Addr:      port,
		Handler:   mw.ResponseTimeMiddleware(mw.SecurityHeaders(mw.Cors(mux))),
		TLSConfig: tlsConfig,
	}

	fmt.Print("Server listening on port ", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
