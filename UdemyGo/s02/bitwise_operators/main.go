package main

import "fmt"

func main() {

	var x uint8
	x = 2

	fmt.Printf("%8d %#8o %#8x %08b\n", x, x, x, x)

	//  binary shit left operation
	x = x << 1
	fmt.Printf("%8d %#8o %#8x %08b\n", x, x, x, x)

	// binary right shift operation
	x = x >> 2
	fmt.Printf("%8d %#8o %#8x %08b\n", x, x, x, x)

	// binary OR
	x = 4 | 2
	fmt.Printf("%8d %#8o %#8x %08b\n", x, x, x, x)

	// binary AND
	x = 4 & 2
	fmt.Printf("%8d %#8o %#8x %08b\n", x, x, x, x)

	// x = ^4  // compiler error: constant -5 overflow
	y := ^4
	fmt.Printf("%8d %#8o %#8x %08b\n", y, y, y, y)

}
