package main

import (
	"fmt"
	"time"
)

func main() {
	go waitAndSay("World")
	fmt.Println("Hello")
	time.Sleep(3 * time.Second)

	go func(s string) {
		time.Sleep(2 * time.Second)
		fmt.Println(s)
	}("World")

	time.Sleep(3 * time.Second)

	word := "Hello"

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println(word)
	}()

	word = "this is changed now"

	time.Sleep(3 * time.Second)
}

func waitAndSay(s string) {
	time.Sleep(2 * time.Second)
	fmt.Println(s)
}
