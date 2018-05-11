package main

import "fmt"

func main() {

	s := "こんにちは世界"

	fmt.Printf("% x\n\n", s)

	r := []rune(s)
	fmt.Println(r)
	fmt.Printf("%x\n", r)

	fmt.Println(string(r))
	fmt.Printf("%q\n", r)
	fmt.Printf("%q\n", s)
}
