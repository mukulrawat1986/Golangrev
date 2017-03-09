package main

import (
	"fmt"
	"math/rand"
	"time"
)

var scMapping = map[string]int{
	"James": 5,
	"Kevin": 10,
	"Rahul": 9,
}

func findSC(name, server string, c chan int) {
	// simulate searching
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)

	// return security clearance from map
	c <- scMapping[name]
}

func main() {

	rand.Seed(time.Now().UnixNano())

	c1 := make(chan int)
	c2 := make(chan int)

	name := "James"

	go findSC(name, "Server 1", c1)
	go findSC(name, "Server 2", c2)

	select {
	case sc := <-c2:
		fmt.Println(name, " has a security clearance of ", sc, " found in server 1")

	case sc := <-c2:
		fmt.Println(name, " has a security clearance of ", sc, " found in server 2")

	case <-time.After(5 * time.Second):
		fmt.Println("Search timed out!!!")
	}
}
