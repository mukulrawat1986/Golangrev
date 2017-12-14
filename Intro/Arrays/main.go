package main

import (
	"fmt"
)

func main() {
	words := [...]string{"the", "quick",
		"brown", "fox",
		"jumps", "over",
		"the", "lazy", "dog"}

	fmt.Printf("%s\n", words[2])
}
