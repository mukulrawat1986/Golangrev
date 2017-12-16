package main

import "fmt"

func emit(c chan string) {
	words := []string{"the", "quick",
		"brown", "fox"}

	for _, word := range words {
		c <- word
	}

	close(c)
}

func main() {
	wordChannel := make(chan string)

	go emit(wordChannel)

	// ranging over a channel means that receive
	// everything from this channel one by one
	// until the channel is closed.
	for word := range wordChannel {
		fmt.Printf("%s ", word)
	}
	fmt.Printf("\n")
}
