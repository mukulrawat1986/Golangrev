package main

import "testing"
import "os"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	// check if deck has correct number of cards
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16. but got %d", len(d))
	}

	// make sure first card is an ace of spades
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades but got %v", d[0])
	}

	// make sure that the last card is four of clubs
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {

	// delete any file with the name _decktesting
	os.Remove("_decktesting")

	// create a new deck
	deck := newDeck()

	// save to file
	deck.saveToFile("_decktesting")

	// load deck from disk
	loadedDeck := newDeckFromFile("_decktesting")

	// assert if deck is correct
	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	// remove the temporary file
	os.Remove("_decktesting")
}
