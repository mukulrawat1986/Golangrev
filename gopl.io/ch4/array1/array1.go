package main

import "fmt"

func main() {

	// declare an array of 3 integers
	var a [3]int

	// print the first element
	fmt.Println(a[0])

	// print the last element
	fmt.Println(a[len(a)-1])

	// print the indices and elements
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// print the elements only
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	// array literal
	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}

	for i, v := range q {
		fmt.Printf("%d %d\n", i, v)
	}

	for i, v := range r {
		fmt.Printf("%d %d\n", i, v)
	}
}
