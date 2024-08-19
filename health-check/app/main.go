package main

import (
	"fmt"
	"net/http"
	"os"
)

var counter = 0

func main() {
	port := os.Getenv("APP_PORT")
	fmt.Println("Server is running on port", port)
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/health", HealthCheck)
	http.ListenAndServe(":"+port, nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	counter++
	if counter > 5 {
		w.WriteHeader(500)
		fmt.Fprint(w, "Internal Server Error")
	} else {
		w.WriteHeader(201)
		fmt.Fprint(w, "Healthy")
	}
}