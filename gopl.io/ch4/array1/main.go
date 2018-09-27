// Array examples
package main

import "fmt"

func main() {

	// array declaration
	var a [10]int

	// array initialization
	for i := 0; i < len(a); i++ {
		a[i] = 0
	}

	// printing the array
	fmt.Println(a[0])
	fmt.Println(a[1])

	// print the indices and element
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// print the elements only
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	// use an array literal to intitialize an arry
	q := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("%#v\n", q)

	// another form of initialization using an array literal
	r := [...]int{99: -1}
	for i := 0; i < len(r); i++ {
		fmt.Printf("%d ", r[i])
	}
	fmt.Println()
}
