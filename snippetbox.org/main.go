package main

import (
	"log"
	"net/http"
	"time"
)

func timeHandler(format string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm))
	}
}

func main() {

	mux := http.NewServeMux()

	mux.Handle("/time/rfc1123", timeHandler(time.RFC1123))
	mux.Handle("/time/rfc3339", timeHandler(time.RFC3339))

	log.Println("Listening.....")
	http.ListenAndServe(":3000", mux)
}
