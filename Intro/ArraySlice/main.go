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
	words = append(words, "The")
	words = append(words, "quick")
	words = append(words, "brown")
	words = append(words, "fox")
	printer(words)
}
