package main

import (
	"fmt"

	"github.com/c81singh/cards/deck"
)

func main() {
	cards := deck.NewDeckFromFile("test.txt")
	cards.Print()
	fmt.Println("")
	sCards := cards.Shuffle()
	sCards.Print()

}
