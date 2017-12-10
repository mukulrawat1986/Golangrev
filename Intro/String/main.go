package main

import (
	"fmt"
)

func main() {
	atoz := "the quick brown fox jumps over the lazy dog"
	fmt.Printf("%s\n", atoz[15:])

	// traversing a string rune by rune
	// range when applied to a string does utf-8 decoding
	// implicitly
	for i, r := range atoz {
		fmt.Printf("%d %c\n", i, r)
	}

	// get the length of the string
	// note: len returns the number of bytes in a string
	fmt.Printf("%d\n", len(atoz))

}
