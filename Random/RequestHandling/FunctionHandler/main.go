package main

import (
	"net/http"
	"time"
	"log"
)

func timeHandler(format string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm))
	}
	return http.HandlerFunc(fn)
}

func main() {

	// create a new servemux
	mux := http.NewServeMux()

	// create a handler
	th := timeHandler(time.RFC1123)

	// add it to servemux
	mux.HandleFunc("/time", th)

	log.Println("Listening......")
	log.Fatal(http.ListenAndServe(":3000", mux))
}