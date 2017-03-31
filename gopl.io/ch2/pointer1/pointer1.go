package main

import "fmt"

func main() {
	// declare and initialize a variable
	x := 10

	// declare a pointer and set it to point towards x
	p := &x

	// print out the type of p, the address it stores and value it points to
	fmt.Printf("%T\n", p)
	fmt.Printf("%x\n", p)
	fmt.Printf("%d\n", *p)

	// change the value of the variable pointed to by p
	*p = 12

	// print out the value of x, and note the change
	fmt.Printf("%d\n", x)
}
