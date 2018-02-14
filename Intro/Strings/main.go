package main

import "fmt"

func main() {
	atoz := "the quick brown fox jumps over the lazy dog\n"

	// going through a string rune by rune
	for _, r := range atoz {
		fmt.Printf("%c\n", r)
	}

	// going through a string rune by rune
	s := "Hello 世界"
	for i, r := range s {
		fmt.Printf("%d %c\n", i, r)
	}

	// number of bytes in a string
	fmt.Printf("Total bytes in %s = %d\n", s, len(s))

}
