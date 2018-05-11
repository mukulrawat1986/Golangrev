package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello世界"
	fmt.Println("The number of bytes in ", s, ": ", len(s))
	fmt.Println("The number of characters in ", s, ": ", utf8.RuneCountInString(s))
}
