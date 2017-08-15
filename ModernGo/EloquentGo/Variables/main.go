package main

import "fmt"

var x unint8 = 2

func main() {
	// x := 2
	// var x uint8 = 2
	fmt.Println(x * 10)

	fmt.Println(compareNumbers(23, 12))

	switch x {
	case 3:
		fmt.Println("I am 3")
	case 2:
		fmt.Println("I am 2")
	case 4:
		fmt.Println("I am 4")
	default:
		fmt.Print("Can't guess me")
	}

	if i := 2; i < 1 {
		fmt.Println(i)
	}

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	// a while in go
	var i int
	for i <= 10 {
		fmt.Println(i)
		i++
	}
}

func compareNumbers(i1, i2 int) (bool, int) {
	switch {
	case i1 > i2:
		fmt.Println("first number is greater than second number")
		return false, i1 - i2

	case i2 > i1:
		fmt.Println("second number is greater than first number.")
		return false, i2 - i1

	default:
		fmt.Print("numbers are equal")
		return true, 0
	}
}
