package main

import "fmt"

func main() {

	atoz := "the quick brown fox jumps over the lazy dog"

	var vowels, consonants int

	for _, r := range atoz {

		// switching over a variable
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			vowels += 1
		default:
			consonants += 1
		}
	}

	fmt.Printf("Vowels: %d; Consonants: %d\n", vowels, consonants)
}
