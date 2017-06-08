// a nil channel will not block when you try to receive from it

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func reader(ch chan int) {
	t := time.NewTimer(10 * time.Second)
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
	stopper := time.NewTimer(2 * time.Second)
	restarter := time.NewTimer(5 * time.Second)

	// save the initiail state of the channel
	savedCh := ch

	for {
		select {
		case ch <- rand.Intn(42):
		case <-stopper.C:
			ch = nil
		case <-restarter.C:
			ch = savedCh
		}
	}
}

func main() {
	ch := make(chan int)
	go reader(ch)
	go writer(ch)

	time.Sleep(20 * time.Second)
}
