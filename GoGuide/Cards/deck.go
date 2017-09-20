package main

import (
	"fmt"
	"strings"
)

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

// Create and return a list of playing cards.
// Essentially an array of strings.
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// create a hand of cards
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// save a list of cards to a file on the local machine
func saveToFile() {

}

// helper function to convert a deck to a string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
