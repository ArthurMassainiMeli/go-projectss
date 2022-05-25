package main

import d "cards-project/deck_features"

func main() {
	deck := d.NewDeck()

	// deal a hand our of deck_features
	hand, _ := deck.DealHand(3)
	hand.Print()

	// save a deck_features in a deck_file
	deck.SaveToFile("my_deck")

	// retrieve a deck_features in a deck_file
	deckTwo := d.NewDeckFromFile("my_deck")
	deckTwo.Print()

	// shuffle a deck_features and print it
	deckTwo.Shuffle()
	deckTwo.Print()
}
