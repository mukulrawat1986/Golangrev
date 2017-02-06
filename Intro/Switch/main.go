package main

import "fmt"

func main() {

	atoz := "the quick brown fox jumps over the lazy dog"

	vowels := 0
	consonants := 0
	zeds := 0

	for _, r := range atoz {
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			vowels++
		case 'z':
			zeds++
			fallthrough
		default:
			consonants++
		}
	}

	fmt.Printf("Vowels: %d; Consonants: %d (Zeds: %d)\n", vowels, consonants, zeds)
}
