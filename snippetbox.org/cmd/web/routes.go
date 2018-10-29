package main

import "net/http"

func (app *App) Routes() *http.ServeMux {
	// Declare a ServeMux and define the routes
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/snippet", app.ShowSnippet)
	mux.HandleFunc("/snippet/new", app.NewSnippet)

	// use app.StaticDir field as location of the static file directory
	fileServer := http.FileServer(http.Dir(app.StaticDir))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// return the serve mux
	return mux
}
