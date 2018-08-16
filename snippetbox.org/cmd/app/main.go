package main

import (
	"log"
	"net/http"
)

func main() {
	// create a new servemux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
