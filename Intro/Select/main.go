package main

import (
	"fmt"
	"time"
)

func emit(wordChannel chan string, done chan bool) {

	defer close(wordChannel)

	words := []string{"The", "quick", "brown",
		"fox"}

	var i int

	t := time.NewTimer(3 * time.Second)

	for {
		select {
		case wordChannel <- words[i]:
			i++
			if i == len(words) {
				i = 0
			}
		case <-done:
			done <- true
			return

		case <-t.C:
			fmt.Printf("\nTimed out\n")
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
