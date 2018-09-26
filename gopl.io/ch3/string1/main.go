// play with strings
package main

import "fmt"

func main() {
	// The len function returns the numbeer of bytes (not runes) in a string.
	// Index operation s[i] retrieves the i-th byte of string.
	// The ith byte of string is not necessarily the ith character of a string
	// because the UTF-8 encoding of non-ASCII code points require two or more bytes.
	s := "hello world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])
	fmt.Printf("%c %c\n", s[0], s[7])

	s = "こんにちは世界"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])
	fmt.Printf("%c %c\n", s[0], s[7])
}
