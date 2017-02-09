package main

import "fmt"

func printer(w []string) {
	for _, word := range w {
		fmt.Printf("%s", word)
	}

	fmt.Printf("\n")
}

func main() {
	// slice
	words := []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}
	printer(words[2:4])
	words[2] = "blue"
	printer(words[2:4])
}
