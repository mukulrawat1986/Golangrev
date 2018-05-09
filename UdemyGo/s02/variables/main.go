package main

import "fmt"

func main() {
	// default declaration of a variable to a zero value
	var i int
	fmt.Println("i = ", i)

	// declaration of a variable
	j := 2
	fmt.Printf("%d - %T\n", j, j)
}
