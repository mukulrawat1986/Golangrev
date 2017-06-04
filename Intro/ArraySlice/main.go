package main

import "fmt"

func main() {
	// an array of strings of unknown number of values
	words := [...]string{"the", "quick", "brown", "fox"}

	fmt.Printf("%s\n", words[2])
}
