package main

import "fmt"

func whatIsThis(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Printf("It's a string %s\n", v)
	case uint32:
		fmt.Printf("It's an unsigned 32-bit integer %d\n", v)
	default:
		fmt.Printf("Don't know %v\n", v)
	}
}

func main() {

	// empty interface
	// an interface with no methods in it, it satisfied everything
	// Gives us an unknown type that's determined at run time.
	whatIsThis(uint32(42))
	whatIsThis("Hello")
	whatIsThis([...]string{"A", "B", "C"})
}
