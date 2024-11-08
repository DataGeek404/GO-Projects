package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Handle requests to the root URL
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "My name is James Muchiri, This is my first server ,brave!!")
	})

	// Start the server on port 8080
	fmt.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
