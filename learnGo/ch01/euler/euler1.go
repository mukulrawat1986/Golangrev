// calculates sum of all multiples of 3 and 5 less than MAX value
package main

import (
	"fmt"
)

const MAX = 1000

func main() {
	work := make(chan int, MAX)
	result := make(chan int)

	// create channel of multiples of 3 and 5
	// concurrently using 	goroutine
	go func() {
		for i := 1; i < MAX; i++ {
			if (i%3) == 0 || (i%5) == 0 {
				work <- i
			}
		}
		close(work)
	}()

	// concurrently sum up work and put result
	// in channel result
	go func() {
		r := 0
		for i := range work {
			r += i
		}
		result <- r
	}()

	// wait for result
	fmt.Println("Total: ", <-result)
}
