package main

import "fmt"

func main() {
	// array with unknown number of values of string
	words := [...]string{"the", "quick", "brown", "fox"}

	fmt.Printf("%s\n", words[2])
}
