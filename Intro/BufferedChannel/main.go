package main

import (
	"fmt"
	"math/rand"
	"time"
)

func work() {
	fmt.Printf("[")
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	fmt.Printf("]")
}

func worker(sema chan bool) {
	// only start work when we receive something
	<-sema

	work()

	// once works is done we will send something to the channel
	sema <- true
}

func main() {
	sema := make(chan bool, 20)

	for i := 0; i < 1000; i++ {
		go worker(sema)
	}

	for i := 0; i < cap(sema); i++ {
		sema <- true
	}

	time.Sleep(30 * time.Second)
}
