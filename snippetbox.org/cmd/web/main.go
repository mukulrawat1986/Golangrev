package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/snippet", ShowSnippet)
	mux.HandleFunc("/snippet/new", NewSnippet)

	// create a file server which serves files out of the "./ui/static" directory
	// The path given to http.Dir is relative to our project repository root
	fileServer := http.FileServer(http.Dir("./ui/static"))

	// register fileServer as handler for all URL paths that start with "/static/"
	// For matching path, we strip "/static" prefix before the request reaches the server
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
