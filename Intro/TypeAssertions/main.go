package main

import "fmt"

func whatIsThis(i interface{}) {
	switch i.(type) {
	case string:
		s := i.(string)
		fmt.Printf("It's a string %s\n", s)
	case uint32:
		u := i.(uint32)
		fmt.Printf("It's an unsigned 32-bit integer %d\n", u)
	default:
		fmt.Printf("Don't know\n")
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
