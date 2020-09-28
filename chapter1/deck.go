package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Suit of a playing card
type Suit string

var suits = []Suit{
	"Spades",
	"Hearts",
	"Clubs",
	"Diamonds",
}

// FaceValue is the face value of a playing card
type FaceValue string

var faceValues = []FaceValue{
	"Ace",
	"Two",
	"Three",
	"Four",
	"Five",
	"Six",
	"Seven",
	"Eight",
	"Nine",
	"Ten",
	"Jack",
	"Queen",
	"King",
}

// PlayingCard represents a typical playing card with Suit and FaceValue
type PlayingCard struct {
	faceValue FaceValue
	suit      Suit
}

// Deck of 52 playing cards
type Deck struct {
	cards [52]PlayingCard
	index int
}

// NewDeck generates a new Deck of 52 PlayingCards
func NewDeck() *Deck {
	var deck Deck
	for i, suit := range suits {
		for j, faceValue := range faceValues {
			var card PlayingCard
			card.suit = suit
			card.faceValue = faceValue
			deck.cards[i*len(faceValues)+j] = card
		}
	}
	return &deck
}

// Shuffle randomizes the order of the PlayingCards in the Deck
// The algorithm is the Fisher–Yates shuffle, see https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle
// It chooses a random card to place at each index from cards not already chosen, liking pulling cards from a hat
func (deck *Deck) Shuffle() {
	deck.index = 0
	for i := range deck.cards {
		swap := rand.Intn(len(deck.cards) - i)
		deck.cards[i], deck.cards[swap + i] = deck.cards[swap + i], deck.cards[i]
	}
}

// DealOneCard returns a PlayingCard off the top of the Deck
func (deck *Deck) DealOneCard() PlayingCard {
	card := deck.cards[deck.index]
	deck.index++
	return card
}

// String representation of a PlayingCard
func (card PlayingCard) String() string {
	return fmt.Sprintf("%v of %v", card.faceValue, card.suit)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var deck = NewDeck()
	deck.Shuffle()
	for range deck.cards {
		fmt.Printf("Dealt the %v\n", deck.DealOneCard().String())
	}
}
