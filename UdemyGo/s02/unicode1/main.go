package main

import "fmt"

func main() {
	ascii := 'a'
	unicode := '平'
	newline := '\n'

	fmt.Printf("%d %[1]c %[1]q %[1]T\n", ascii)
	fmt.Printf("%d %[1]c %[1]q %[1]T\n", unicode)
	fmt.Printf("%d %[1]c %[1]q %[1]T\n", newline)
}
