package main

import "fmt"

func main() {

	// infinite loop
	for {
		fmt.Printf("Hello World\n")
	}

	// using counter to run a loop fixed amount of time
	var counter int
	counter = 0

	for counter < 10 {
		fmt.Printf("Hello World\n")
		counter += 1
	}

	// shorthand notation
	for i := 0; i < 10; i++ {
		fmt.Printf("Hello World\n")
	}

	// using two variables in the loop
	for i, j := 0, 1; i < 10; i, j = i+1, j*2 {
		fmt.Printf("%d %d\n", i, j)
	}

	// using a boolean
	var stop boolean

	i := 0

	for !stop {
		fmt.Printf("Hello, World!\n")
		i += 1
		if i == 10 {
			stop = !stop
		}
	}
}
