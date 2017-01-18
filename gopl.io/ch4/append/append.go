// showing how the append function works

package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	fmt.Println(a)
	a = appendInt(a, 6)
	a = appendInt(a, 8)
	a = appendInt(a, 9)
	fmt.Println(a)

	var x, y []int
	for i := 0; i < 10; i++ {
		x = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1

	if zlen <= cap(x) {
		// There is room to grow
		// extend the slice
		z = x[:zlen]
	} else {
		// There is insufficient space. Allocate a new array
		// Grow by doubling for amortized linear complexity
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}
