package main

import "fmt"

func emit(wordChannel chan string, done chan bool) {
	words := []string{"The", "quick", "brown",
		"fox"}

	var i int

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
