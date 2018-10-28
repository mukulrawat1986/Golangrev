package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

// Home is a method against *App
func (app *App) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Because the Home handler function is now a method against App, it can access
	// its fields. So we can build the paths to the HTML template files
	files := []string{
		filepath.Join(app.HTMLDir, "base.html"),
		filepath.Join(app.HTMLDir, "home.page.html"),
	}

	// use template.ParseFiles() to create a template set
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Use ExecuteTemplate() to execute the base template and write its content
	// to http.ResponseWriter
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server error", 500)
	}

}

// ShowSnippet is a method against *App
func (app *App) ShowSnippet(w http.ResponseWriter, r *http.Request) {
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

// NewSnippet is a method against *App
func (app *App) NewSnippet(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Display a new snippet form...."))
}
