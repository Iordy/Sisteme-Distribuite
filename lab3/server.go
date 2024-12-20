package main

import (
	"fmt"
	"net/http"
)

func writer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!") // Respond with "Hello, world!"
}

func server() {

	http.HandleFunc("/hello", writer)

	http.HandleFunc("/", writer)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}

}
