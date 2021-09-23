package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Deck []string

func (deck Deck) Print() {
	for _, card := range deck {
		fmt.Println(card)
	}
}

func NewDeck() Deck {
	deck := Deck{}
	suits := allSuits()
	cards := allCards()
	for _, suit := range suits {
		for _, card := range cards {
			deck = append(deck, card+" of "+suit)
		}
	}
	return deck
}

func (deck Deck) shuffle() {
	rand.Seed(time.Now().UTC().UnixNano())
	l := len(deck)
	for i := range deck {
		newi := rand.Intn(l - 1)
		deck[i], deck[newi] = deck[newi], deck[i]
	}
}

func (deck Deck) Draw() string {
	return deck[0]
}

func deal(deck Deck, size int) (Deck, Deck) {
	return deck[:size], deck[size:]
}

func (deck Deck) toString() string {
	return strings.Join([]string(deck), ",")
}

func allCards() []string {
	return []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
}

func allSuits() []string {
	return []string{"Hearts", "Diamonds", "Spades", "Clubs"}
}
