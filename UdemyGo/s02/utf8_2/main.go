package main

import "fmt"

func main() {

	s := "こんにちは世界"

	fmt.Println("The number of bytes in ", s, ": ", len(s))

	// iterate through the runes or the unicode code points and print out
	// the code points unicode value and its printed representation also
	// print out its byte position
	for i, v := range s {
		fmt.Printf("%#U  %d starts at byte position %d\n", v, v, i)
	}

	fmt.Println([]byte(s))

	fmt.Println()
	s = "界ABCD"
	fmt.Println("The number of bytes in ", s, ": ", len(s))
	for i, v := range s {
		fmt.Printf("%#U %d starts at byte position %d\n", v, v, i)
	}
}
