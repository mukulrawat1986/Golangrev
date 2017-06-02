package main

import "fmt"

func main() {

	atoz := "the quick brown fox jumps over the lazy dog\n"

	// take substrings
	fmt.Printf("%s\n", atoz[0:9])
	fmt.Printf("%s\n", atoz[:9])
	fmt.Printf("%s\n", atoz[15:19])
	fmt.Printf("%s\n", atoz[15:])

	// going through a string rune by rune
	for i, r := range atoz {
		fmt.Printf("%d %c\n", i, r)
	}

	// length of string
	// len gives the number of bytes in the string
	fmt.Println(len(atoz))
}
