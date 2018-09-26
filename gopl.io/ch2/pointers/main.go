package main

import "fmt"

func main() {
	returnPointer()
	playIncr()
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

func playIncr() {
	// By passing a pointer argument to a function, it becomes possible for
	// the function to update the variable that was indirectly passed.
	v := 1
	fmt.Println(incr(&v), v)
	fmt.Println(incr(&v), v)
	fmt.Println(incr(&v), v)
	fmt.Println(incr(&v), v)
	fmt.Println(incr(&v), v)
	fmt.Println(incr(&v), v)
	fmt.Println(incr(&v), v)
	fmt.Println(incr(&v), v)
}

func incr(p *int) int {
	*p++ // increments what p points to, does not change p
	return *p
}
