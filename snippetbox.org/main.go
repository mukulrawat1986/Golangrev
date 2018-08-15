package main

import (
	"log"
	"net/http"
)

// Home function writes a plain-test "Hello from Snippetbox" message as the HTTP response body.
func Home(w http.ResponseWriter, r *http.Request) {

	// use r.URL.Path to check whether the request URL path
	// exactly matches "/". If it does not, return a 404.
	if r.URL.Path != "/" {
		w.WriteHeader(404)
		w.Write([]byte("Not Found"))
		return
	}

	w.Write([]byte("Hello from Snippetbox"))
}

func main() {
	// create a new servemux
	mux := http.NewServeMux()

	// register the Home function as the handler for the "/"
	// url pattern
	mux.HandleFunc("/", Home)

	// start our server and listen on port 4000
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
