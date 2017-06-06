package main

import "fmt"

func emit(c chan string) {

	words := []string{"the", "quick", "brown", "fox"}
	for _, word := range words {
		// pass the string on the channel
		c <- word
	}
	// close the channel
	close(c)
}

func main() {

	wordChannel := make(chan string)

	// start a goroutine
	go emit(wordChannel)

	// receive the words coming from the channel
	// when we range over a channel, we receive everything on this channel
	// successfully, until the channel is closed
	for word := range wordChannel {
		fmt.Printf("%s ", word)
	}
	fmt.Printf("\n")
}
