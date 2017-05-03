package main

import "fmt"

// increment declares count as variable whose value is always
// an integer.
// go: noinline
func increment(inc int) {

	// increment the "value of" inc
	inc++
	fmt.Println("Inc: ", inc, &inc)
}

func main() {

	// Declare variable of type int with a value of 10.
	count := 10

	// Display the "value of" and "address of" count.
	fmt.Println("Before: ", count, &count)

	// Pass the "value of" the variable count
	increment(count)

	fmt.Println("After: ", count, &count)
}
