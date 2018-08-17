package main

import (
	"net/http"
)

// Routes return the servemux holding all the routes in the application
func (app *App) Routes() *http.ServeMux {
	// create a new servemux
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/snippet", app.ShowSnippet)
	mux.HandleFunc("/snippet/new", app.NewSnippet)

	// create a file server which serves files out of "./web/static"
	// The path is relative to our project repository root
	fileServer := http.FileServer(http.Dir(app.StaticDir))

	// use mux.Handle to register fileServer as a handler for all url
	// paths that start with "/static/".
	// For matching paths, we strip the /static/ prefix before the
	// request reaches the filleServer handler
	mux.Handle("/static/", http.StripPrefix("/static",
		fileServer))

	return mux
}
