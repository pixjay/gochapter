package main

import (
	"fmt"
	"math/rand"
	"time"
	"io"
	"log"
	"net/http"
	"sync"
	"strconv"
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

func (deck Deck) cardsRemaining() int {
	return len(deck.cards) - deck.index
}

// String representation of a PlayingCard
func (card PlayingCard) String() string {
	return fmt.Sprintf("%v of %v", card.faceValue, card.suit)
}

type player struct {
	id int
	name string
	hand []PlayingCard
}

type game struct {
	id int
	deck Deck
	players []*player
	mux sync.Mutex
}

func (game *game) fprintf(w io.Writer) {
	fmt.Fprintf(w, "Game %v (cards remaining, %v):\n", game.id, game.deck.cardsRemaining())
	for _, player := range game.players {
		fmt.Fprintf(w, " - %v (id %v) has a %v card hand\n", player.name, player.id, len(player.hand))
	}
}

var gamesMux sync.Mutex
var games[]*game

func newGame() *game {
	newGame := new(game)
	gamesMux.Lock()
	games = append(games, newGame)
	newGame.id = len(games) - 1
	gamesMux.Unlock()
	return newGame
}

func newPlayer(game *game, name string) *player {
	newPlayer := new(player)
	newPlayer.name = name
	game.mux.Lock()
	game.players = append(game.players, newPlayer)
	newPlayer.id = len(game.players) - 1
	game.mux.Unlock()
	return newPlayer
}

func rootHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Dealer, the web server")
}

func gameHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		for _, game := range games {
			game.fprintf(w)
		}
	case "POST": 
		game := newGame()
		fmt.Fprintf(w, "New game created with id: %v", game.id)
	}
}

func gamePlayerHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	gameID, _ := strconv.ParseInt(req.Form.Get("gameID"), 10, 64)
	game := games[gameID]
	switch req.Method {
	case "GET":
		for _, player := range game.players {
			fmt.Fprintf(w, "%v has a %v card hand", player.name, len(player.hand))
		}
	case "POST":
		name := req.Form.Get("name")
		player := newPlayer(game, name)
		fmt.Fprintf(w, "New player created with id: %v", player.id)
	}
}

func gamePlayerDealHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	gameID, _ := strconv.ParseInt(req.Form.Get("gameID"), 10, 64)
	game := games[gameID]
	playerID, _ := strconv.ParseInt(req.Form.Get("playerID"), 10, 64)
	player := game.players[playerID]
	game.mux.Lock()
	player.hand = append(player.hand, game.deck.DealOneCard())
	game.mux.Unlock()
	game.fprintf(w)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/game", gameHandler)
	http.HandleFunc("/game/player", gamePlayerHandler)
	http.HandleFunc("/game/player/deal", gamePlayerDealHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
