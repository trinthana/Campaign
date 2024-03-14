package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Message struct to read the input and write the output
type Message struct {
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Go Application")
	})

	http.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var msg Message

			// Decode the incoming JSON payload
			err := json.NewDecoder(r.Body).Decode(&msg)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			// Set the content type to JSON
			w.Header().Set("Content-Type", "application/json")

			// Encode and send the message back as JSON
			err = json.NewEncoder(w).Encode(msg)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		} else {
			// If the method is not POST, return an error
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server starting on port 443...")
	log.Fatal(http.ListenAndServe(":443", nil))
}
