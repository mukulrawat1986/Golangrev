package main

import (
	"log"
	"net/http"
)

// Home function writes a plan-text message as the HTTP response body
func Home(w http.ResponseWriter, r *http.Request) {
	// use request.URL.Path to check whether the request URL path exactly matches '/'
	if r.URL.Path != "/" {
		w.WriteHeader(404)
		w.Write([]byte("Not Found"))
		return
	}

	w.Write([]byte("Hello from Snippetbox"))
}

func main() {
	// Initialize a new serve mux
	mux := http.NewServeMux()

	// register our home function as the handler for the "/" url pattern
	mux.HandleFunc("/", Home)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
