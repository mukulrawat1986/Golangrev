package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {

	// Define command line flags for the network address and location of
	// the static files directory.
	addr := flag.String("addr", ":4000", "HTTP Network Address")
	staticDir := flag.String("static-dir", "./web/static", "Path to static assets")
	htmlDir := flag.String("html-dir", "./web/html", "Path to HTML Template")

	flag.Parse()

	// Initialize a new instance of App containing the dependencies
	app := &App{
		HTMLDIR: *htmlDir,
	}

	// create a new servemux
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/snippet", app.ShowSnippet)
	mux.HandleFunc("/snippet/new", app.NewSnippet)

	// create a file server which serves files out of "./web/static"
	// The path is relative to our project repository root
	fileServer := http.FileServer(http.Dir(*staticDir))

	// use mux.Handle to register fileServer as a handler for all url
	// paths that start with "/static/".
	// For matching paths, we strip the /static/ prefix before the
	// request reaches the filleServer handler
	mux.Handle("/static/", http.StripPrefix("/static",
		fileServer))

	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
