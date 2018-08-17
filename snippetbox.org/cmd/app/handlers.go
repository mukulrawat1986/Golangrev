package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// Home handler
func (app *App) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.NotFound(w)
		return
	}

	app.RenderHTML(w, "home.page.html")
}

// ShowSnippet Handler function
func (app *App) ShowSnippet(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id parameter from the query string
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	// if id can't be converted to an integer or the value is less than 1
	// we return a 404
	if err != nil || id < 1 {
		app.NotFound(w)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet (ID %d)...", id)
}

// NewSnippet Handler function
func (app *App) NewSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display the new snippet form....\n"))
}
