// buffered channels allows us to send even when there is no one
// to receive upto a fixed capacity

package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

var (
	running int64
)

func work() {
	atomic.AddInt64(&running, 1)
	fmt.Printf("[%d", running)
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	atomic.AddInt64(&running, -1)
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
