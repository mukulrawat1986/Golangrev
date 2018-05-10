package main

import "fmt"

func main() {
	// basic boolean values
	var b1 = true
	var b2 = false

	// logical operations
	fmt.Println(b1 || b2)
	fmt.Println(b1 && b2)
	fmt.Println(b1 && !b2)
	fmt.Println(true && !false)

}
