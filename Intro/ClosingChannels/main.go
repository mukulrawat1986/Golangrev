// coordination by closing closing channels

package main

import (
	"fmt"
	"time"
)

func printer(msg string, goCh chan bool) {
	<-goCh
	fmt.Printf("%s\n", msg)
}

func main() {

	goCh := make(chan bool)

	for i := 0; i < 10; i++ {
		go printer(fmt.Sprintf("printer: %d", i), goCh)
	}

	time.Sleep(5 * time.Second)

	// by closing the goChannel, all goroutines will get a zero value
	// from the channel
	close(goCh)

	time.Sleep(5 * time.Second)

}
