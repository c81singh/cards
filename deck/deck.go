package deck

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

type Deck []string

type Suit int

const (
	Diamonds Suit = iota
	Cubs
	Hearts
	Spades
)

func (d Deck) Print() {
	for i, v := range d {
		fmt.Printf("%d -> %s\n", i, v)
	}
}

func Deal(d Deck, handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

func (d Deck) ToString() string {
	return strings.Join(d, ",")
}

func (d Deck) SaveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.ToString()), 0755)
}

func NewDeckFromFile(filename string) Deck {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to open the file %s.\n ERROR: %s", filename, err)
		return Deck{}
	}
	s := string(f)
	ss := strings.Split(s, ",")
	return ss
}

func (d Deck) ToString2() []byte {
	var byty []byte
	for _, v := range d {
		byty = append(byty, []byte(v)...)
	}
	return byty
}

func (d Deck) ToString3() []byte {
	var byty []byte
	var newline string = "\n"
	for _, v := range d {
		for _, bv := range v {
			byty = append(byty, byte(bv))
		}
		byty = append(byty, []byte(newline)...)
	}
	return byty
}

func (d Deck) Shuffle() Deck {
	rand.Seed(time.Now().UnixMicro())
	for i, _ := range d {
		rn := rand.Intn(len(d))
		d[i], d[rn] = d[rn], d[i]
	}
	return d
}

func NewDeck() Deck {
	var deck []string = []string{}
	suits := []string{"Diamonds", "Cubs", "Hearts", "Spades"}
	cValues := []string{
		"Ace",
		"King",
		"Queen",
		"Jack",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten"}

	for _, suit := range suits {
		for _, v := range cValues {
			deck = append(deck, v+" of "+suit)
		}
	}

	return deck

}
