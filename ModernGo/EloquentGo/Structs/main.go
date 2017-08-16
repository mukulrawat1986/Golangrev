package main

import "fmt"

type person struct {
	name    string
	age     int
	address string
}

func main() {

	// declare a struct using literals
	jason := person{
		name:    "Json S",
		age:     38,
		address: "Germany",
	}

	fmt.Println(jason)
}
