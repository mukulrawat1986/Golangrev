package main

import "fmt"

func printer(w []string) {
	for _, word := range w {
		fmt.Printf("%s", word)
	}

	fmt.Printf("\n")

	w[2] = "blue"
}

func main() {
	// slice
	words := []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}
	fmt.Printf("%s\n", words[2])
	fmt.Printf("%d\n", len(words))
	printer(words)
	printer(words)
}
