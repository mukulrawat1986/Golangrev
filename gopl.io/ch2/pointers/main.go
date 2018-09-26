package main

import "fmt"

func main() {
	returnPointer()
}

func returnPointer() {
	// Function returns the address of a local variable and the local variable
	// remains in existence even after the call has returned.
	p := f()
	fmt.Println(*p, p)
	*p++
	fmt.Println(*p, p)
	*p++
	fmt.Println(*p, p)
	*p++
	fmt.Println(*p, p)
	*p++
	fmt.Println(*p, p)
	*p++
	fmt.Println(*p, p)
}

func f() *int {
	v := 1
	return &v
}
