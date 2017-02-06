package main

import "fmt"

func main() {
	var counter int

	counter = 0

	for counter < 10 {
		fmt.Printf("Hello World!\n")
		counter++
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("Another Hello World!\n")
	}

	// declaring more than one variable in loop
	for i, j := 0, 1; i < 10; i, j = i+1, j+2 {
		fmt.Printf("%d And Another Hello World!!!\n", j)
	}

	var stop bool

	i := 0

	for !stop {
		fmt.Printf("Helllllll\n")
		i++
		if i == 10 {
			stop = true
		}
	}

}
