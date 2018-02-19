package main

import "fmt"

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
	words = append(words, "Quick")
	words = append(words, "Brown")
	words = append(words, "Fox")
	fmt.Printf("%d %d\n", len(words), cap(words))
	printer(words)
	words = append(words, "Jumps")
	fmt.Printf("%d %d\n", len(words), cap(words))

	// copying a slice
	newWords := make([]string, 4)
	copy(newWords, words)
	printer(newWords)
}
