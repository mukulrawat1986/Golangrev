package main

import (
	"net/http"
	"time"
	"log"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is: " + tm))
}

func main() {

	// create a new servemux
	mux := http.NewServeMux()

	// add it to servemux
	mux.HandleFunc("/time", timeHandler)

	log.Println("Listening......")
	log.Fatal(http.ListenAndServe(":3000", mux))
}