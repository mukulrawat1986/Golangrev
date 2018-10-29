package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// Home is a method against *App
func (app *App) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.NotFound(w)
		return
	}

	app.RenderHTML(w, "home.page.html")
}

// ShowSnippet is a method against *App
func (app *App) ShowSnippet(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id parameter from the query string and try to
	// convert it to an integer using the strconv.Atoi() function
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	// if id < 1 or couldn't be converted to an integer return a 404
	if err != nil || id < 1 {
		app.NotFound(w)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet (ID %d)\n", id)
}

// NewSnippet is a method against *App
func (app *App) NewSnippet(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Display a new snippet form...."))
}
