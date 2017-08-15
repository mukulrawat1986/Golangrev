package main

import "fmt"

func main() {
	fmt.Println(compareNumbers(2, 2))
	fmt.Println(compareNumbers(3, 6))
	fmt.Println(compareNumbers(8, 4))
}

func compareNumbers(i1, i2 int) (bool, int) {
	if i1 > i2 {
		return false, i1 - i2
	} else if i2 > i1 {
		return false, i2 - i1
	}
	return true, 0
}
