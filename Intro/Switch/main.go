package main

import (
	"fmt"
	"os"
)

func main() {
	n, err := fmt.Print("Hello, World!")

	fmt.Printf("\n")

	switch {
	case err != nil:
		os.Exit(1)

	case n == 0:
		fmt.Printf("No bytes output")

	case n != 13:
		fmt.Printf("Wrong number of characters")

	default:
		fmt.Printf("OK!")
	}

	fmt.Printf("\n")
}
