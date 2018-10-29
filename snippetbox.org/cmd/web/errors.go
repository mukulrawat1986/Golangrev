package main

import (
	"log"
	"net/http"
	"runtime/debug"
)

// ServerError helper writes an error message and a stack trace to the log, then sends a generic
// 500 Internal Server Error response to the user
func (app *App) ServerError(w http.ResponseWriter, err error) {
	log.Printf("%s\n%s", err.Error(), debug.Stack())
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

// ClientError helper sends a specific status code and description to the user
func (app *App) ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// NotFound helper is a wrapper around Client Error and returns 404 Not Found response to the User
func (app *App) NotFound(w http.ResponseWriter) {
	app.ClientError(w, http.StatusNotFound)
}
