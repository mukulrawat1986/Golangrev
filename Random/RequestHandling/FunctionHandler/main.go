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

	// convert the timeHandler function to a HandlerFunc type
	th := http.HandlerFunc(timeHandler)

	// add it to servemux
	mux.Handle("/time", th)

	log.Println("Listening......")
	log.Fatal(http.ListenAndServe(":3000", mux))
}