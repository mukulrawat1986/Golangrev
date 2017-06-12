package main

import "fmt"

func main() {

	// create a variable by shorcircuiting and not using var
	message := "The answer to life is %d\n"
	answer := 42

	fmt.Printf(message, answer)
}
