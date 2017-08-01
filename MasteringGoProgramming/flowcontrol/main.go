package main

import "fmt"

func main() {
	// if statement
	i := 10
	if i < 0 {
		fmt.Println("i is a negative number")
	} else if i == 0 {
		fmt.Println("i is zero")
	} else {
		fmt.Println("i is a positive number")
	}
}
