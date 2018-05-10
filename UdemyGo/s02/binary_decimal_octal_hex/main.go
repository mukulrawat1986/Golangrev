package main

import "fmt"

func main() {

	a := byte('A')

	// showing print format specifier for decimal, octal, hex
	// and binary.
	fmt.Printf("%d %o %x %b \n", a, a, a, a)

	// using only one argument with multiple format specifiers
	fmt.Printf("%d %[1]o %[1]x %[1]b", a)
}
