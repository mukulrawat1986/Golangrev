package main

import (
	"net/http"
	"time"
	"log"
)

func timeHandler(format string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm)
	}
	return http.HandlerFunc(fn)
}

func main() {
	var format string = time.RFC1123
	th := timeHandler(format)

	// we use http.Handle
	http.Handle("/time", th)

	log.Println("Listening.....")
	// and pass nil as the handler to ListenAndServe
	http.ListenAndServe(":3000", nil)
}
