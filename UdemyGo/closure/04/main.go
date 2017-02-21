package main

import "fmt"

func main() {
	x := 0
	increment := func() int {
		x++
		return x
	}

	// if we change value of x over here, the results would
	// be different

	fmt.Println(increment())
	fmt.Println(increment())
}
