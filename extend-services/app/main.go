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
	http.ListenAndServe(":"+port, nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	mode := os.Getenv("MODE")
	response := "Hello " + mode
	fmt.Fprint(w, response)
}