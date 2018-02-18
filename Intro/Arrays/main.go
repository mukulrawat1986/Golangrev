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

	words := [4]string{"the", "quick", "brown", "fox"}

	printer(words)
	printer(words)
}
