package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

func ShowSnippet(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id parameter from the query string and try to
	// convert it to an integer using the strconv.Atoi() function
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	// if id < 1 or couldn't be converted to an integer return a 404
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet (ID %d)\n", id)
}

func NewSnippet(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Display a new snippet form...."))
}
