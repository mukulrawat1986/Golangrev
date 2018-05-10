package main

import "fmt"

func main() {

	// width of integer, width of binary and width of binary with leading zeroes.
	fmt.Printf("%3d %8b %08b\n", 2, 2, 2)
	fmt.Printf("%3d %8b %08b\n\n", -2, -2, -2)

	fmt.Printf("%3d %08b\n", 3, 3)
	fmt.Printf("%3d %08b\n\n", -3, -3)

	fmt.Printf("%3d %08b\n", 10, 10)
	fmt.Printf("%3d %08b\n\n", -10, -10)
}
