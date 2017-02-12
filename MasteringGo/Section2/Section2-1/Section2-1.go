package Utility

import "fmt"

// SayHello is to say hello to the world
func SayHello() {
	fmt.Println("Hello!!")
}

// Say is to print out a string and an integer
func Say(s string, i int) {
	fmt.Println(s, i)
}

func addSubtractMultiply(a, b int) (addition, subtraction, multiplication int) {
	// or return a+b, a-b, a*b
	addition = a + b
	subtraction = a - b
	multiplication = a * b
	return
}
