// emitting on a single channel from two goroutines, and communicating
// between the goroutines on a different channel
package main

import "fmt"

func emit(wordChan chan string, commChan chan bool, id int) {

	words := []string{"the", "quick", "brown", "fox"}

	for _, word := range words {
		// pass the string on the channel
		wordChan <- word
	}

	if id == 0 {
		commChan <- true
		<-commChan
	} else {
		<-commChan
		commChan <- true
	}

	close(wordChan)

}

func main() {

	// channel to pass the string on
	wordChan := make(chan string)

	// channel to communicate between goroutines
	commChan := make(chan bool)

	go emit(wordChan, commChan, 0)
	go emit(wordChan, commChan, 1)

	for word := range wordChan {
		fmt.Printf("%s ", word)
	}

	fmt.Printf("\n")

}
