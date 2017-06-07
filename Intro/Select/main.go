package main

import "fmt"

func emit(wordChan chan string, done chan bool) {
	// slice of string
	words := []string{"the", "quick", "brown", "fox"}

	var i int

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
		}
	}
}

func main() {
	wordCh := make(chan string)
	doneCh := make(chan bool)

	go emit(wordCh, doneCh)

	for i := 0; i < 100; i++ {
		fmt.Printf("%s ", <-wordCh)
	}

	doneCh <- true
	<-doneCh
}
