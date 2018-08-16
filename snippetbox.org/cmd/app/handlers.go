package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// Home handler
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox\n"))
}

// ShowSnippet Handler function
func ShowSnippet(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id parameter from the query string
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	// if id can't be converted to an integer or the value is less than 1
	// we return a 404
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet (ID %d)...", id)
}

// NewSnippet Handler function
func NewSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display the new snippet form....\n"))
}
