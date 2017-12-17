package main

import (
	"fmt"
	"math/rand"
)

func makeRandoms(c chan int) {
	for {
		c <- rand.Intn(1000)
	}
}

func main() {
	randoms := make(chan int)
	go makeRandoms(randoms)

	for n := range randoms {
		fmt.Printf("%d ", n)
	}

}
