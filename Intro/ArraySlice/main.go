package main

import "fmt"

func printer(w [4]string) {
	for _, word := range w {
		fmt.Printf("%s", word)
	}
	fmt.Printf("\n")
}

func main() {
	// an array of strings
	words := [4]string{"the", "quick", "brown", "fox"}
	printer(words)
}
