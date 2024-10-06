package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Struct to handle the POST request body
type EchoRequest struct {
	Message string `json:"message"`
}

// GET /greet handler
func greetHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, welcome to the Go API!"))
}

// POST /echo handler
func echoHandler(w http.ResponseWriter, r *http.Request) {
	var echoReq EchoRequest
	err := json.NewDecoder(r.Body).Decode(&echoReq)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	response := fmt.Sprintf("You said: %s", echoReq.Message)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func main() {
	// Define the routes
	http.HandleFunc("/hello", greetHandler)
	http.HandleFunc("/echo", echoHandler)

	// Start the server on port 8080
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
