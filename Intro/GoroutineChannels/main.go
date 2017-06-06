package main

import "fmt"

func makeID(idChan chan int) {
	var id int
	for {
		idChan <- id
		id++
	}
}

func main() {
	idChan := make(chan int)

	go makeID(idChan)

	// generate a id
	fmt.Printf("%d\n", <-idChan)

	fmt.Printf("%d\n", <-idChan)
}
