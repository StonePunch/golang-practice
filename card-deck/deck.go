package main

import (
	"fmt"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	deck := deck{}

	cardSuits := []string{
		"Spades",
		"Diamonds",
		"Hearts",
		"Clubs",
	}
	cardValues := []string{
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

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			deck = append(deck, fmt.Sprintf(`%s of %s`, value, suit))
		}
	}

	return deck
}

func deal(deck deck, handSize int) (hand, remainder deck) {
	return deck[:handSize], deck[handSize:]
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

func (d deck) toString() string {
	const separator string = ","

	return strings.Join([]string(d), separator)
}

func (d deck) saveToFile(filename string) error {
	deckByteSlice := []byte(d.toString())

	return os.WriteFile(filename, deckByteSlice, 0666)
}