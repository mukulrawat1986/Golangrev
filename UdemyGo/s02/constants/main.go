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

	// introducing iota which is a constant generator.
	// It starts with 0 and keeps on incrementing by one.
	const (
		a2 = iota + 1
		b2
		c2
		d2 = a2 * 10
		e2
		f2 = iota
	)

	fmt.Println(a2, b2, c2, d2, e2, f2)
}
