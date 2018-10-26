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

	// some variables create from slices of the months slice
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)
	fmt.Println(summer)

	// common elements
	for _, s := range summer {
		for _, q := range Q {
			if s == q
		}
	}
}
