package main

import (
	"github.com/c81singh/cards/deck"
)

func main() {
	cards := deck.NewDeckFromFile("test.txt")
	cards.Print()
}
