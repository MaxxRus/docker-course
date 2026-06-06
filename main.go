package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Go web server!")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health", healthHandler)

	fmt.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
