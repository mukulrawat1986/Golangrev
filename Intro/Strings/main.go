package main

import "fmt"

func main() {

	atoz := "The quick brown fox jumps over the lazy dog\n"

	// take substrings of the string
	fmt.Printf("%s\n", atoz[:9])
	fmt.Printf("%s\n", atoz[15:19])
	fmt.Printf("%s\n", atoz[15:])

	// String is made of runes which are unicode characters

	// we can go through a string rune by rune
	// we do it by using a for loop
	for i, r := range atoz {
		fmt.Printf("%d %c\n", i, r)
	}

	// if we don't wanna use the index
	for _, r := range atoz {
		fmt.Printf("%c\n", r)
	}

	// len gives us the number of bytes in the string
	// When we are using ASCII text, that's the length of the string
	fmt.Printf("Length of the string: %d\n", len(atoz))

}
