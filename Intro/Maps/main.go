package main

import "fmt"

func main() {
	// create a map using a literal
	dayMonths := map[string]int{
		"Jan": 31,
		"Feb": 28,
		"Mar": 31,
		"Apr": 30,
		"May": 31,
		"Jun": 30,
		"Jul": 31,
		"Aug": 31,
		"Sep": 30,
		"Oct": 31,
		"Nov": 30,
		"Dec": 31,
	}

	for month, day := range dayMonths {
		fmt.Printf("%s %d\n", month, day)
	}
}
