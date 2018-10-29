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
		StaticDir: *staticDir,
	}

	// Again, we dereference the addr variable and use it as the network address
	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, app.Routes())
	log.Fatal(err)
}
