package main

import (
	"fmt"
)

func printer(w [9]string) {
	for _, word := range w {
		fmt.Printf("%s", word)
	}
	fmt.Printf("\n")
	w[2] = "blue"
}

func main() {
	words := [...]string{"the", "quick",
		"brown", "fox",
		"jumps", "over",
		"the", "lazy", "dog"}

	fmt.Printf("%s\n", words[2])
	printer(words)
	printer(words)
}
