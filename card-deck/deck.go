package main

import (
	"fmt"
	"os"
	"strings"
)

const separator string = ","

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

func newDeckFromFile(filename string) deck {
	deckByteSlice, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	deckString := string(deckByteSlice)

	return strings.Split(deckString, separator)
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
	return strings.Join([]string(d), separator)
}

func (d deck) saveToFile(filename string) error {
	deckByteSlice := []byte(d.toString())

	return os.WriteFile(filename, deckByteSlice, 0666)
}
