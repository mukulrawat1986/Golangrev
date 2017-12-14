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
	words := make([]string, 4)
	words[0] = "The"
	words[1] = "quick"
	words[2] = "brown"
	words[3] = "fox"
	printer(words)
}
