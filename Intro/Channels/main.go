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

	for {
		word, ok := <-wordChannel
		if !ok {
			fmt.Printf("Channel has been closed\n")
			break
		} else {
			fmt.Printf("%s\n", word)
		}
	}
}
