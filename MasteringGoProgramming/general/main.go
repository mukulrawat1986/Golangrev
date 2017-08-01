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
}
