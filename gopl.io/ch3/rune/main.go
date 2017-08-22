package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "hello, 世界"
	// print out number of bytes
	fmt.Println(len(s))

	// print out the number of runes
	fmt.Println(utf8.RuneCountInString(s))

	// process utf-8 characters
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\t%d\n",i, r, size)
		i+=size
	}

	// Go's range operation applies utf-8 decoding automatically
	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	// count number of runes without using utf8 package
	size := 0
	for _, _ = range s {
		size++
	}
	fmt.Println(size)
}
