package main

import "fmt"

func main() {

	months := []string{
		1: "January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}
	for i, v := range months {
		fmt.Printf("%d\t%s\n", i, v)
	}
}
