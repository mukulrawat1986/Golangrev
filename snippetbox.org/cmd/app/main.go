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
		HTMLDir:   *htmlDir,
		StaticDir: *staticDir,
	}

	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, app.Routes())
	log.Fatal(err)
}
