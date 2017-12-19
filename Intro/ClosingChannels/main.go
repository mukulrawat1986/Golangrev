package main

import (
	"fmt"
	"time"
)

func printer(msg string, stopCh chan bool) {
	for {
		select {
		case <-stopCh:
			fmt.Printf("Work stopped\n")
			return
		default:
			fmt.Printf("%s\n", msg)
		}
	}
}

func main() {
	stopCh := make(chan bool)

	for i := 0; i < 10; i++ {
		go printer(fmt.Sprintf("printer: %d", i), stopCh)
	}

	time.Sleep(1 * time.Second)
	close(stopCh)
	time.Sleep(3 * time.Second)
}
