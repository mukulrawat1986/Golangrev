package main

import "fmt"

func printer(w [4]string) {
	for _, word := range w {
		fmt.Printf("%s", word)
	}

	fmt.Printf("\n")

	w[2] = "blue"
}

func main() {
	// Array with unknown number of strings
	words := [...]string{"the", "quick", "brown", "fox"}
	fmt.Printf("%s\n", words[2])
	printer(words)
	printer(words)
}
