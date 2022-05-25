package deck_features

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := NewDeck()
	deckSize := len(deck)

	if deckSize != 52 {
		t.Errorf("Expected the length of 52, but got %v", deckSize)
	}

	if deck[0] != "Ace of Diamonds" {
		t.Errorf("Expected the first card to be Ace of Diamonds, but got %v", deck[0])
	}

	if deck[deckSize-1] != "King of Spades" {
		t.Errorf("Expected the last card to be King of Spades, but got %v", deck[deckSize-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	deckTestingFileName := "_deck-testing"
	deck := NewDeck()
	deck.SaveToFile(deckTestingFileName)

	loadedDeck := NewDeckFromFile(deckTestingFileName)

	if len(loadedDeck) != 52 {
		t.Errorf("Expected the length of 52, but got %v", len(loadedDeck))
	}

	err := os.Remove(deckTestingFileName)
	if err != nil {
		os.Exit(1)
	}
}
