package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var startedAt = time.Now()

// Main sets up a server that will respond to any request with the string
// "Hello Full Cycle!".
func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/config-map", ConfigMap)
	http.HandleFunc("/secret", Secret)
	http.HandleFunc("/healthz", Healthz)
	http.ListenAndServe(":8080", nil)
}

// Hello responds to a request with a friendly message using the
// NAME and AGE environment variables.
func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	fmt.Fprintf(w, "Hello, I'm %s and I'm %s years old!", name, age)
}

// ConfigMap reads a file named family.txt from a directory named myFamily and writes
// the contents of that file to the HTTP response writer.
func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("myFamily/family.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	fmt.Fprintf(w, "My Family: %s", string(data))
}

// Secret writes the USER and PASSWORD environment variables to the HTTP response
// writer. This is a terrible idea and should never be done in a real application.
func Secret(w http.ResponseWriter, r *http.Request) {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	fmt.Fprintf(w, "User: %s\nPassword: %s", user, password)
}

// Healthz is a health check function that returns 200 if the server has been running
// for between 10 and 30 seconds, and 500 otherwise.
func Healthz(w http.ResponseWriter, r *http.Request) {
	duration := time.Since(startedAt)
	if duration.Seconds() < 10 {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Duration: %v", duration.Seconds())))
	} else {
		w.WriteHeader(200)
		w.Write([]byte("OK"))
	}
}
