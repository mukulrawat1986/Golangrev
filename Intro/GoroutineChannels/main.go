package main

import "fmt"

func emit(c chan string) {

	words := []string{"the", "quick", "brown", "fox"}
	for _, word := range words {
		// pass the string on the channel
		c <- word
	}
	// close the channel
	//	close(c)
}

func main() {

	wordChannel := make(chan string)

	// start a goroutine
	go emit(wordChannel)
	go emit(wordChannel)

	// we range over the channel till the channel is open
	// if the channel is not closed in the goroutine and we are
	// ranging over it a deadlock happens
	for word := range wordChannel {
		fmt.Printf("%s ", word)
	}

	fmt.Printf("\n")
}
