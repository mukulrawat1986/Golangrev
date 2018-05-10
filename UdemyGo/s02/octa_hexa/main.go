package main

import "fmt"

func main() {

	d := 125
	fmt.Printf("%8d %#8o %#8x \t %08b\n", d, d, d, d)

	oc := 0175
	fmt.Printf("%8d %#8o %#8x \t %08b\n", oc, oc, oc, oc)

	x := 0x7d
	fmt.Printf("%8d %#8o %#8x \t %08b\n", x, x, x, x)

	// b := 01111101 // binary literals are not supported in Go
}
