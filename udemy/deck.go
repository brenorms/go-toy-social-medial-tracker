package main

import "fmt"

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

func (deck Deck) Draw() string {
	return deck[0]
}

func allCards() []string {
	return []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
}

func allSuits() []string {
	return []string{"Hearts", "Diamonds", "Spades", "Clubs"}
}
