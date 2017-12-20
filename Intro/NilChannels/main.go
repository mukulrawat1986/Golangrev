package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Nil channles don't block when we try to receive from it

func reader(ch chan int) {
	t := time.NewTimer(3 * time.Second)
	for {
		select {
		case i := <-ch:
			fmt.Printf("%d\n", i)

		case <-t.C:
			ch = nil
		}
	}
}

func writer(ch chan int) {
	t := time.NewTimer(2 * time.Second)
	
	for {
		select {
		case ch <- rand.Intn(42):
		case <-t.C:
			fmt.Printf("stop sending\n")
			ch = nil
		}
	}
}

func main() {
	ch := make(chan int)

	go reader(ch)
	go writer(ch)

	time.Sleep(10 * time.Second)
}
