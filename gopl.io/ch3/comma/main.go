// comma inserts commas in a non-negative decimal integer string
package main

import (
	"fmt"
)

func main() {
	fmt.Println(comma("1"))
	fmt.Println(comma("12"))
	fmt.Println(comma("213"))
	fmt.Println(comma("3451"))
	fmt.Println(comma("23451"))
	fmt.Println(comma("234543"))
	fmt.Println(comma("1345678"))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
