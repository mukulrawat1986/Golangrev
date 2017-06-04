package main

import "fmt"

func printer(w []string) {
	for _, word := range w {
		fmt.Printf("%s", word)
	}
	fmt.Printf("\n")
}

func main() {
	// using make to create a slice of string and appending
	// items to it
	words := make([]string, 0, 4)
	fmt.Printf("%d %d\n", len(words), cap(words))
	words = append(words, "The")
	words = append(words, "quick")
	words = append(words, "brown")
	words = append(words, "fox")
	fmt.Printf("%d %d\n", len(words), cap(words))
	printer(words)
	words = append(words, "jumps")
	fmt.Printf("%d %d\n", len(words), cap(words))

	newWords := make([]string, 4)
	copy(newWords, words)
	newWords[2] = "blue"
	printer(newWords)
}
