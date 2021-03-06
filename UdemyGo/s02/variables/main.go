package main

import "fmt"

func main() {
	// default declaration of a variable to a zero value
	var i int
	fmt.Println("i = ", i)

	// declaration of a variable
	j := 2
	fmt.Printf("%d - %T\n", j, j)

	// another way of declaration and initialization, not much in use
	// use uint8
	var k1 uint8 = 20
	fmt.Printf("%d - %T\n", k1, k1)

	var k2 uint16 = 30

	// kind of type cohersion
	k2 = uint16(k1)
	fmt.Printf("%d - %T\n", k2, k2)
}
