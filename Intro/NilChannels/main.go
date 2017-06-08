// a nil channel will not block when you try to receive from it

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func reader(ch chan int) {
	t := time.NewTimer(3 * time.Second)
	for {
		select {
		// when ch becomes nil, this particular case is ignored by select
		case i := <-ch:
			fmt.Printf("%d\n", i)
		case <-t.C:
			ch = nil
		}
	}
}

func writer(ch chan int) {
	for {
		ch <- rand.Intn(42)
	}
}

func main() {
	ch := make(chan int)
	go reader(ch)
	go writer(ch)

	time.Sleep(10 * time.Second)
}
