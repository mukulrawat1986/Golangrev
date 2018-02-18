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

	words := []string{"the", "quick", "brown", "fox"}

	printer(words)
	printer(words)
}
