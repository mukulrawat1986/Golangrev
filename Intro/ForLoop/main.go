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

}
