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

	word := <-wordChannel
	fmt.Printf("%s\n", word)
	word = <-wordChannel
	fmt.Printf("%s\n", word)
	word = <-wordChannel
	fmt.Printf("%s\n", word)
	word = <-wordChannel
	fmt.Printf("%s\n", word)

	// now if we try to receive from the channel again
	word = <-wordChannel
	fmt.Printf("%s\n", word)

}
