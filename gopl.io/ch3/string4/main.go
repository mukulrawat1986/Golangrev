package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	s := "Hello, 世痌"
	// number of bytes in the string
	fmt.Println("len of ", s, ": ", len(s))
	// number of runes(characters or unicode code points) in the string
	fmt.Println("Number of runes in ", s, ": ", utf8.RuneCountInString(s))

	fmt.Println()

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	fmt.Println()
	fmt.Println("Decoding UTF-8 using the range loop")

	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	// count runes in strings without using utf8.RuneCountInString
	count := 0
	for range s {
		count++
	}
	fmt.Println("The number of runes in ", s, ": ", count)

	s = "プログラム"
	fmt.Printf("% x\n", s)
	r := []rune(s)
	fmt.Printf("% x\n", r)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c\n", r[i])
	}

	fmt.Println(string(r))

	fmt.Println(string(65))
}
