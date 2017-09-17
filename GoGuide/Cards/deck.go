package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// print loops through a deck of cards
// and print out the value of every single card
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
