package main

import "fmt"

func printer(w [4]string) {
	for _, word := range w {
		fmt.Printf("%s", word)
	}
	fmt.Printf("\n")

	// since its an array, the change won't be reflected
	w[2] = "blue"
}

func printer1(w []string) {
	for _, word := range w {
		fmt.Printf("%s", word)
	}
	fmt.Printf("\n")

	// since its a slice, the change will be reflected
	w[2] = "blue"
}

func main() {
	// declare and initialize an array of strings
	words := [...]string{"the", "quick", "brown", "fox"}

	// fmt.Printf("%s\n", words[2])

	printer(words)
	printer(words)

	// we use slice, which is a window to an underlying
	// array
	words1 := []string{"the", "quick", "brown", "fox",
		"jumps", "over", "the", "lazy", "dog"}

	fmt.Printf("The length of string is %d\n", len(words1))

	printer1(words1)
	printer1(words1)
}
