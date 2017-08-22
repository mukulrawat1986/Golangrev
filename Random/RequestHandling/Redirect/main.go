package main

import (
	"net/http"
	"log"
)

func main() {
	// create a new servemux
	mux := http.NewServeMux()

	// create a redirect handler
	rh := http.RedirectHandler("http://example.org", 307)
	mux.Handle("/foo", rh)

	log.Println("Listening.....")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
