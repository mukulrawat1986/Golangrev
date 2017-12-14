package main

import (
	"fmt"
)

func printer(w []string) {
	for _, word := range w {
		fmt.Printf("%s", word)
	}
	fmt.Printf("\n")
	w[2] = "blue"
}

func main() {
	words := []string{"the", "quick",
		"brown", "fox",
		"jumps", "over",
		"the", "lazy", "dog"}

	fmt.Printf("Length of the array is: ")
	fmt.Printf("%d\n", len(words))

	fmt.Printf("%s\n", words[2])
	printer(words)
	printer(words)
}
