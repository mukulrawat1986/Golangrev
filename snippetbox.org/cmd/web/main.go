package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {

	// Define command line flags for the network address and location of the
	// static files directory.
	addr := flag.String("addr", ":4000", "HTTP network address")
	htmlDir := flag.String("html-dir", "./ui/html", "Path to HTML templates")
	staticDir := flag.String("static-dir", "./ui/static", "Path to static assets")

	// Parse the command-line flags
	flag.Parse()

	// Initialize a new instance of App containing the dependencies
	app := &App{
		HTMLDir: *htmlDir,
	}

	// use App object's methods as the handler functions
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/snippet", app.ShowSnippet)
	mux.HandleFunc("/snippet/new", app.NewSnippet)

	// create a file server which serves files out of the "./ui/static" directory
	// The path given to http.Dir is relative to our project repository root
	// The value returned from the flag.String() function is a pointer to a string, so we need
	// to derefernce it
	fileServer := http.FileServer(http.Dir(*staticDir))

	// register fileServer as handler for all URL paths that start with "/static/"
	// For matching path, we strip "/static" prefix before the request reaches the server
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Again, we dereference the addr variable and use it as the network address
	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
