package main

import "fmt"

func main() {

	var k1 uint8
	k1 = 20
	// constants are variables whose value cannot be changed.
	const pi = 3.14159

	k1 /= 2

	// variables are statically typed.
	// Their values can change, but their types can never be changed.
	// k1 = "changing type is not allowed" // compile error

	fmt.Println(25, pi, k1)

	// declaring and assigning multiple constants at once
	// the constants get the same value as their preceeding constant
	// unless its changed.
	const (
		a = 2
		b
		c
		d = a * 10
		e
		f
	)

	fmt.Println(a, b, c, d, e, f)
}
