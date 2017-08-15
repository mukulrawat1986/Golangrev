package main

import "fmt"

func main() {
	// array is a collection of values of same datatype

	// create an array of 5 integers
	var a [5]int

	// create an array of 5 integers and assign them values
	var b [5]int = [5]int{1, 2, 3, 4, 5}

	//  short way to declare arrays
	c := [5]int{1, 2, 3, 4, 5}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	mySlice := []int{1, 2, 3, 4, 5} // []int{}
	fmt.Println(mySlice)

	// appending to a slice
	mySlice = append(mySlice, 6, 7, 8)

	fmt.Println(mySlice)
}
