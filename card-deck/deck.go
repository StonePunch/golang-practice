package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
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

func newDeckFromFile(filename string) (deck, error) {
	deckByteSlice, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return deck{}, err
	}

	deckString := string(deckByteSlice)

	return strings.Split(deckString, separator), nil
}

// Deal func will return a deck with the first x cards from the inputted
// deck corresponding to the handSize. A deck with the leftover cards
// will also be returned.
//
// If the handSize is larger that the length of the deck, the full deck
// will be returned along with an empty deck as the remainder
func deal(d deck, handSize int) (hand, remainder deck) {
	if handSize > len(d) {
		return d, deck{}
	}

	return d[:handSize], d[handSize:]
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

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(source)

	for i := range d {
		pos := randGen.Intn(len(d) - 1)

		d[i], d[pos] = d[pos], d[i]
	}
}
