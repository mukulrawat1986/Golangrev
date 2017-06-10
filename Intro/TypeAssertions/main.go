package main

import "fmt"

func whatIsThis(i interface{}) {
	fmt.Printf("%T\n", i)
}

func main() {

	// empty interface
	// an interface with no methods in it, it satisfied everything
	// Gives us an unknown type that's determined at run time.
	whatIsThis(42)
}
