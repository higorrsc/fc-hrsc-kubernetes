package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Main sets up a server that will respond to any request with the string
// "Hello Full Cycle!".
func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/config-map", ConfigMap)
	http.ListenAndServe(":8080", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	fmt.Fprintf(w, "Hello, I'm %s and I'm %s years old!", name, age)
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("myFamily/family.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	fmt.Fprintf(w, "My Family: %s", string(data))
}
