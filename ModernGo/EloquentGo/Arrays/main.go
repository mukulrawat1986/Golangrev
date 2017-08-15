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

	// a slice of length and capacity of 5
	s := make([]int, 5)
	s[0], s[1], s[2], s[3], s[4] = 1, 2, 3, 4, 5 // capacity: cap(s), length: len(s)
	fmt.Println(s)

	s1 := s[1:3]
	fmt.Println(s1[:cap(s1)])

	// shallow copy
	s3 := s1
	s3[0] = 10000
	fmt.Println(s3, s1)

	// copying slices - deep copy
	s2 := make([]int, 2)
	copy(s2, s1)
	fmt.Println(s2)
}
