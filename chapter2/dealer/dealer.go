package main

import (
	"time"
	"io"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"strconv"
	"fmt"
	"github.com/pixjay/gochapter/chapter2/deck"
)

type player struct {
	id int
	name string
	hand []deck.PlayingCard
}

type game struct {
	id int
	deck deck.Deck
	players []*player
	mux sync.Mutex
}

func (game *game) fprintf(w io.Writer) {
	fmt.Fprintf(w, "Game %v (cards remaining, %v):\n", game.id, game.deck.CardsRemaining())
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
