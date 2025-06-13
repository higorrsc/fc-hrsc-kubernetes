package main

import (
	"fmt"
	"net/http"
	"os"
)

// Main sets up a server that will respond to any request with the string
// "Hello Full Cycle!".
func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8080", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	fmt.Fprintf(w, "Hello, I'm %s and I'm %s years old!", name, age)
}
