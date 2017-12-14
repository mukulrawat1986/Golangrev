package main

import (
	"fmt"
)

func printer(w []string) {
	for _, word := range w {
		fmt.Printf("%s", word)
	}
	fmt.Printf("\n")
}

func main() {
	words := make([]string, 0, 4)
	fmt.Printf("%d %d\n", len(words), cap(words))
	words = append(words, "The")
	words = append(words, "quick")
	words = append(words, "brown")
	words = append(words, "fox")
	fmt.Printf("%d %d\n", len(words), cap(words))
	printer(words)

	newWords := make([]string, 4)
	copy(newWords, words)
	printer(newWords)
	newWords[2] = "blue"
	printer(words)

	newWords1 := words[0:4]
	copy(newWords1, words)
	printer(newWords1)
	newWords1[2] = "blue"
	printer(words)

}
