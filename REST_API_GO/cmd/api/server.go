package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	mw "restapi/internal/api/middlewares"
	"sync"
)

type Teacher struct {
	ID        int
	FirstName string
	LastName  string
	Class     string
	Subject   string
}

var (
	teachers = make(map[int]Teacher)
	mutex    = &sync.Mutex{}
	nextID   = 1
)

func init() {
	teachers[nextID] = Teacher{
		ID:        nextID,
		FirstName: "John",
		LastName:  "Doe",
		Class:     "9A",
		Subject:   "Math",
	}
	nextID++
	teachers[nextID] = Teacher{
		ID:        nextID,
		FirstName: "Jane",
		LastName:  "Smith",
		Class:     "10A",
		Subject:   "Algebra",
	}
}

func getTeachersHandler(w http.ResponseWriter, r *http.Request) {
	teacherList := make([]Teacher, 0, len(teachers))
	for _, teacher := range teachers {
		teacherList = append(teacherList, teacher)
	}
	response := struct {
		Status string    `json:"status"`
		Count  int       `json:"count"`
		Data   []Teacher `json:"data"`
	}{
		Status: "success",
		Count:  len(teacherList),
		Data:   teacherList,
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
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
		getTeachersHandler(w, r)
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

	//rl := mw.NewRateLimitter(5, time.Minute)
	//hppOptions := mw.HTTPOptions{
	//	CheckQuery:                  true,
	//	CheckBody:                   true,
	//	CheckBodyOnlyForContentType: "application/x-www-form-url-encoded",
	//	Whitelist:                   []string{"sortBy", "sortOrder", "name", "age", "class"},
	//}
	// secureMux := mw.Cors(r1.Middleware(mw.ResponseTimeMiddleware(mw.SecurityHeaders(mw.Compression(mw.Hpp(hppOptions)(mux))))))
	// secureMux := applyMiddlewareHandler(mux, mw.Hpp(hppOptions), mw.Compression, mw.SecurityHeaders, mw.ResponseTimeMiddleware, rl.Middleware, mw.Cors)
	secureMux := mw.SecurityHeaders(mux)
	server := &http.Server{
		Addr:      port,
		Handler:   secureMux,
		TLSConfig: tlsConfig,
	}

	fmt.Print("Server listening on port ", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

type Middleware func(http.Handler) http.Handler

func applyMiddlewareHandler(handler http.Handler, middlewares ...Middleware) http.Handler {

	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler

}
