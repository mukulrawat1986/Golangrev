package main

import "fmt"

func main() {

	cards := newDeck()
	cards.saveToFile("my_cards")

	cardsNew := newDeckFromFile("my_cards")
	for i, card := range cardsNew {
		fmt.Println(i, card)
	}
}
