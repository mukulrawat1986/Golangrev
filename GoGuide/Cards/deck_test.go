package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	// check if deck has correct number of cards
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16. but got %d", len(d))
	}
}
