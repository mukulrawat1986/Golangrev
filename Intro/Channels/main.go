package main

import "fmt"

func emit(c chan string) {
	words := []string{"the", "quick",
		"brown", "fox"}

	for _, word := range words {
		c <- word
	}

	// close(c)
}

func main() {
	wordChannel := make(chan string)

	go emit(wordChannel)
	go emit(wordChannel)

	for word := range wordChannel {
		fmt.Printf("%s ", word)
	}
	fmt.Printf("\n")
}
