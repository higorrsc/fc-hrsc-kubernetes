package main

import "net/http"

// Main sets up a server that will respond to any request with the string
// "Hello Full Cycle!".
func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8080", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello Full Cycle!</h1>"))
}
