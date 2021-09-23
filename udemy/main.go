package main

import (
	"fmt"
)

func main() {
	deck := NewDeck()
	fmt.Println(deck.toString())
	deck.shuffle()
	fmt.Println(deck.toString())
	deck.shuffle()
	fmt.Println(deck.toString())
}
