package main

import (
	"fmt"
	"time"
)

func emit(wordChan chan string, done chan bool) {

	defer close(wordChan)

	// slice of string
	words := []string{"the", "quick", "brown", "fox"}

	var i int
	t := time.NewTimer(3 * time.Second)

	for {
		select {
		// if someone is listening on the wordChan ready to receive
		case wordChan <- words[i]:
			i++
			if i == 4 {
				i = 0
			}

		// if it receives message on done channel
		case <-done:
			done <- true
			return

		// receiving on the timer's Channel
		case <-t.C:
			return
		}
	}
}

func main() {
	wordCh := make(chan string)
	doneCh := make(chan bool)

	go emit(wordCh, doneCh)

	for word := range wordCh {
		fmt.Printf("%s ", word)
	}

}
