package main

import "fmt"

func main() {
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

	fmt.Printf("Days in February :%d\n",
		dayMonths["Feb"])

	days, ok := dayMonths["January"]
	if !ok {
		fmt.Printf("Can't get days for January\n")
	} else {
		fmt.Printf("%d\n", days)
	}

	for month, days := range dayMonths {
		fmt.Printf("%s has %d days\n",
			month, days)
	}

	has31 := 0

	for _, days := range dayMonths {
		if days == 31 {
			has31++
		}
	}

	fmt.Printf("%d months have 31 days\n", has31)

	delete(dayMonths, "Feb")

	for month, days := range dayMonths {
		fmt.Printf("%s month has %d days\n",
			month, days)
	}
}
