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
	words := make([]string, 0)
	words[0] = "the"
	words[1] = "quick"
	words[2] = "brown"
	words[3] = "fox"
	printer(words)
}
