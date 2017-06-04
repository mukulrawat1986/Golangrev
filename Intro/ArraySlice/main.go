package main

import "fmt"

func printer(w []string) {
	for _, word := range w {
		fmt.Printf("%s", word)
	}
	fmt.Printf("\n")
}

func main() {
	// a slice of strings
	words := []string{"the", "quick", "brown", "fox", "jumps",
		"over", "the", "lazy", "dog"}

	// length of slice
	fmt.Printf("%d\n", len(words))

	// slicing a slice
	printer(words[0:2])
	printer(words[:3])
}
