package main

import "fmt"

func main() {
	// create a map
	dayMonths := make(map[string]int)
	dayMonths["Jan"] = 31
	dayMonths["Feb"] = 28
	dayMonths["Mar"] = 31
	dayMonths["Apr"] = 30
	dayMonths["May"] = 31
	dayMonths["Jun"] = 30
	dayMonths["Jul"] = 31
	dayMonths["Aug"] = 31
	dayMonths["Sep"] = 30
	dayMonths["Oct"] = 31
	dayMonths["Nov"] = 30
	dayMonths["Dec"] = 31

	fmt.Printf("Days in February: %d\n", dayMonths["Feb"])

	// accessing a key that does not exist
	// it returns the zero value of the corresponding value
	fmt.Printf("Days in January: %d\n", dayMonths["January"])

	// Distinguishing between error, since key did not exit and the zero value
	days, ok := dayMonths["Januaray"]
	if !ok {
		fmt.Printf("Can't get days for January")
	} else {
		fmt.Printf("%d\n", days)
	}

	// iterating over a map
	for month, days := range dayMonths {
		fmt.Printf("%s has %d days\n", month, days)
	}

	var has31 int

	for _, days := range dayMonths {
		if days == 31 {
			has31 += 1
		}
	}

	fmt.Printf("%d months have 31 days\n", has31)
}
