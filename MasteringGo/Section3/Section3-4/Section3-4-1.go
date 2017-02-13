package main

import "fmt"

func main() {
	c := make(chan bool)
	go waitAndSay(c, "world")
	fmt.Println("Hello")

	// we send a signal to c in order to allow waitAndSay to continue
	c <- true

	// we wait to receive another signal on c before we exit
	<-c
}

func waitAndSay(c chan bool, s string) {
	if b := <-c; b {
		fmt.Println(s)
	}
	c <- true
}
