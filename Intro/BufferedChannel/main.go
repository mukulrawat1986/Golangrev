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
	fmt.Printf("[ %d", running)
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	atomic.AddInt64(&running, -1)
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
