package main

import (
	"log"
	"net/http"
)

func main() {
	// create a new servemux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/snippet", ShowSnippet)
	mux.HandleFunc("/snippet/new", NewSnippet)

	// create a file server which serves files out of "./web/static"
	// The path is relative to our project repository root
	fileServer := http.FileServer(http.Dir("./web/static"))

	// use mux.Handle to register fileServer as a handler for all url
	// paths that start with "/static/".
	// For matching paths, we strip the /static/ prefix before the
	// request reaches the filleServer handler
	mux.Handle("/static/", http.StripPrefix("/static",
		fileServer))

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
