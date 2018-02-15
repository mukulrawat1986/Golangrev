package main

import (
	"fmt"
	"os"
)

func main() {
	n, err := fmt.Printf("Hello, World!\n")

	n = 0

	switch {
	case err != nil:
		os.Exit(1)
	case n == 0:
		fmt.Printf("No bytes output\n")
		fallthrough
	case n != 14:
		fmt.Printf("Wrong number of characters")
	default:
		fmt.Printf("OK!")
	}

	fmt.Printf("\n")
}
