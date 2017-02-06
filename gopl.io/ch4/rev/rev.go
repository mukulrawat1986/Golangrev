// reverse reverses a slice of ints in place

package main

import "fmt"

func main() {
	x := []int{4, 3, 1, 2, 6, 7}
	fmt.Println(x)
	reverse(x)
	fmt.Println(x)

	a := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(a)
	reverse(a[:])
	fmt.Println(a)
}

func reverse(x []int) {
	for i, j := 0, len(x)-1; i < j; i, j = i+1, j-1 {
		x[i], x[j] = x[j], x[i]
	}
}