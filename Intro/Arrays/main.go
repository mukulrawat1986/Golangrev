package main

import "fmt"

func printer(w []string) {
	for _, word := range w {
		fmt.Printf("%s", word)
	}

	fmt.Printf("\n")
}

func main() {
	// using make to create a slice
	words := make([]string, 0, 4)
	words = append(words, "the")
	words = append(words, "quick")
	words = append(words, "brown")
	words = append(words, "fox")
	printer(words)

	newWords := make([]string, 4)

	// create a copy of the slice
	copy(newWords, words)
	newWords[2] = "blue"
	printer(newWords)
	printer(words)

	// create a reference of the slice
	newRef := words[:4]
	newRef[2] = "blue"
	printer(newRef)
	printer(words)
}
