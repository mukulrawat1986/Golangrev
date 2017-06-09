// buffered channels allows us to send even when there is no one
// to receive upto a fixed capacity

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
	// worker will only start working once they recive something on the
	// sema channel
	<-sema

	work()

	// once the work is done, we will put something back in the channel
	sema <- true
}

func main() {
	sema := make(chan bool, 10)

	for i := 0; i < 1000; i++ {
		go worker(sema)
	}

	for i := 0; i < cap(sema); i++ {
		sema <- true
	}

	time.Sleep(30 * time.Second)

}
