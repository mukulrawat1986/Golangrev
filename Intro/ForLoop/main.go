package main

import "fmt"

func main() {
	var stop bool
	var counter int

	for !stop {
		fmt.Printf("Hello, World!\n")
		counter += 1
		if counter == 10 {
			stop = true
		}
	}
}
