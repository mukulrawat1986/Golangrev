package main

import (
	"fmt"
	"os"
)

func main() {
	n, err := fmt.Print("Hello, World!")

	fmt.Printf("\n")

	n = 0

	switch {
	case err != nil:
		os.Exit(1)

	case n == 0:
		fmt.Printf("No bytes output\n")
		fallthrough

	case n != 13:
		fmt.Printf("Wrong number of characters\n")

	default:
		fmt.Printf("OK!\n")
	}

}
