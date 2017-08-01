package main

import "fmt"

func main() {
	// arrays
	var a [2]int
	a[0], a[1] = 4, 5
	fmt.Println("Array a", a, len(a))

	b := [3]int{9, 8, 10}
	fmt.Println("Array b", b, len(b))

	// slice
	c := []int{99, 44}
	c = append(c, 82)
	fmt.Println("Slice c", c, len(c))

	// pointer
	var p1 *int // declare a pointer
	i := 4      // create an integer using type inference
	p1 = &i     // assign the address of the integer to the pointer
	fmt.Println(p1)

	p2 := &i         // create a pointer via type inference
	fmt.Println(*p2) // use the * in order to dereference the pointer and get the original value

}
