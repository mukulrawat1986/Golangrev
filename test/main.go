package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var (
	mu      sync.Mutex
	counter int
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/count/", Counter)

	log.Println("Server listening at :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

func Home(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	counter++
	mu.Unlock()
	w.WriteHeader(200)
	w.Write([]byte("Welcome home\n"))
}

func Counter(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", counter)
	mu.Unlock()
}
