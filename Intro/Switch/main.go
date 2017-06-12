package main

import (
	"fmt"
	"os"
)

func main() {
	n, err := fmt.Printf("Hello, World\n")

	// a general switch statement
	switch {
	case err != nil:
		os.Exit(1)

	case n == 0:
		fmt.Printf("No bytes output\n")

	case n != 14:
		fmt.Printf("Wrong number of characters: %d\n", n)
		fallthrough // making the control fall through to default
	default:
		fmt.Printf("OK!\n")
	}

	atoz := "the quick brown fox jumps over the lazy dog"
	vowels := 0
	consonants := 0
	zeds := 0

	for _, r := range atoz {

		// Switching on a particular variable
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			vowels += 1

		case 'z':
			zeds += 1
			fallthrough

		default:
			consonants += 1
		}
	}

	fmt.Printf("Vowels: %d; Consonants: %d (Zeds: %d)\n", vowels, consonants, zeds)

}
