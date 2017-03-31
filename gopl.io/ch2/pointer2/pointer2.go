package main

import "fmt"

func f() *int {
	v := 1
	return &v
}

func main() {
	x := f()
	fmt.Printf("%T %x %d \n", x, x, *x)
}
