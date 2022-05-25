package deck_features

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Deck []string

func NewDeck() Deck {
	deck := make(Deck, 0, 52)
	cardSuits := []string{"Diamonds", "Clubs", "Hearts", "Spades"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			deck = append(deck, value+" of "+suit)
		}
	}
	return deck
}

func (d *Deck) DealHand(handSize int) (Deck, Deck) {
	*d = (*d)[handSize:]
	return (*d)[:handSize], (*d)[handSize:]
}

func (d Deck) Print() {
	for i, value := range d {
		fmt.Println(i, value)
	}
}

func (d Deck) toString() string {
	return strings.Join(d, ",")
}

func toSlice(s string) Deck {
	return strings.Split(s, ",")
}

func (d Deck) SaveToFile(filename string) {
	err := ioutil.WriteFile("files/"+filename, []byte(d.toString()), 0666)
	if err != nil {
		fmt.Printf("Error saving deck_features\nError: %s\n", err.Error())
		os.Exit(1)
	}

}

func NewDeckFromFile(filename string) Deck {
	byteSlice, err := ioutil.ReadFile("files/" + filename)
	if err != nil {
		fmt.Printf("Error getting deck_features from deck_file\nError: %s\n", err.Error())
		os.Exit(1)
	}
	return toSlice(string(byteSlice))
}

func (d Deck) Shuffle() {
	// generate a new seed to get real random numbers every time
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	deckSize := len(d)

	for i := range d {
		random := r.Intn(deckSize - 1)
		aux := d[i]
		d[i] = d[random]
		d[random] = aux
	}
}

func (d Deck) ShuffleOtherWay() Deck {
	shuffledDeck := make(Deck, 0, 52)
	deckSize := len(d)

	for i := deckSize; i > 0; i-- {
		random := rand.Intn(i)
		shuffledDeck = append(shuffledDeck, d[random])
		removeByIndex(d, random)
	}
	return shuffledDeck
}

func removeByIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func (d Deck) DealRandomHand(qtdOfCards int) []string {
	hand := make([]string, 0, qtdOfCards)
	cardsLeft := 54

	for i := 0; i < qtdOfCards; i++ {
		random := rand.Intn(cardsLeft)
		hand = append(hand, d[random])
		removeByIndex(d, random)
		cardsLeft--
	}
	return hand
}
