package main

import (
	"fmt"
)

func main() {
	deck := NewDeck()
	deck.Print()
	fmt.Println(deck.Draw())
}
