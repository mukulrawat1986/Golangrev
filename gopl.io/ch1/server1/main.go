// server1 is a minimal "echo" server
package main

import (
	"net/http"
	"log"
	"fmt"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the request url r
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf(w, "URL.Path=%q\n", r.URL.Path)
}