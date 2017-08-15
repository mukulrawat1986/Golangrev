package main

import "fmt"

func main() {
	fmt.Println(compareNumbers(2, 2))
	fmt.Println(compareNumbers(3, 6))
	fmt.Println(compareNumbers(8, 4))
}

func compareNumbers(i1, i2 int) (bool, int) {
	if i1 > i2 {
		fmt.Println("First number is greather than second number")
		return false, i1 - i2
	} else if i2 > i1 {
		fmt.Println("Second number is greater than first number")
		return false, i2 - i1
	}
	fmt.Println("Numbers are equal!!!")
	return true, 0
}
