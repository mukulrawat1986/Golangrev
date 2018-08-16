package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Home handler
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Initialize a slice containing the paths to the two files.
	files := []string{
		"./web/html/base.html",
		"./web/html/home.page.html",
	}

	// use the templates.ParseFiles() function to read the file and store the
	// templates in a template set.
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Use the ExecuteTemplate() method to execute the base template and
	// write its content to our ResponseWriter.
	// The last parameter to ExecuteTemplate represents any
	// dynamic data that we want to pass in which for now we will leave as nil.
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
	}
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
