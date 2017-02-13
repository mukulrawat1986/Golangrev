package main

import "fmt"

func main() {
	c := make(chan string)
	go sayHelloMultipleTimes(c, 5)
	for s := range c {
		fmt.Println(s)
	}
}

func sayHelloMultipleTimes(c chan string, n int) {
	for i := 0; i < n; i++ {
		c <- "Hello"
	}
	close(c)
}
