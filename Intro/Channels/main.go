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

	word := <-wordChannel
	fmt.Printf("%s\n", word)
	word = <-wordChannel
	fmt.Printf("%s\n", word)
	word = <-wordChannel
	fmt.Printf("%s\n", word)
	word = <-wordChannel
	fmt.Printf("%s\n", word)
	word, ok := <-wordChannel
	if !ok {
		fmt.Printf("Channel has been closed\n")
	} else {
		fmt.Printf("%s\n", word)
	}
}
