// Array examples
package main

import "fmt"

func main() {

	// array declaration
	var a [3]int

	// array initialization
	for i := 0; i < len(a); i++ {
		a[i] = 0
	}

	// printing the array
	fmt.Println(a[0])
	fmt.Println(a[1])
}
