package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	testCases := []struct {
		name string
		test func(d deck)
	}{
		{
			name: "Deck Length",
			test: func(d deck) {
				if len(d) != 52 {
					t.Errorf("Expected deck length of 52, but was %d", len(d))
				}
			},
		},
		{
			name: "First Card",
			test: func(d deck) {
				if d[0] != "Ace of Spades" {
					t.Errorf("Expected 1st card to be a Ace of Spades, but was %s", d[0])
				}
			},
		},
		{
			name: "Last Card",
			test: func(d deck) {
				if d[len(d)-1] != "King of Clubs" {
					t.Errorf("Expected 52nd card to be a King of Clubs, but was %s", d[len(d)-1])
				}
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.test(d)
		})
	}
}

func TestSaveToFile(t *testing.T) {
	const filename string = "_deckTesting"
	deck := newDeck()

	err := deck.saveToFile(filename)
	if err != nil {
		t.Errorf("Expected deck to be saved, error: %v", err)
	}

	file, err := os.Open(filename)
	if err != nil || file == nil {
		t.Errorf("Expected file named '%s' to exists, error: %v", filename, err)
	}

	// Cleanup file created for testing
	t.Cleanup(func() {
		os.Remove(filename)
	})
}

func TestNewDeckFromFile(t *testing.T) {
	const filename string = "_deckTesting"
	d := newDeck()
	_ = d.saveToFile(filename)

	testCases := []struct {
		name string
		test func(d deck)
	}{
		{
			name: "Successful Load",
			test: func(d deck) {
				loadedDeck, err := newDeckFromFile(filename)
				if err != nil {
					t.Errorf("Expected nil error, but was: %v", err)
				}
				if len(loadedDeck) != 52 {
					t.Errorf("Expected deck length of 52, but was %d", len(loadedDeck))
				}
			},
		},
		{
			name: "Unsuccessful Load",
			test: func(d deck) {
				_, err := newDeckFromFile("")
				if err == nil {
					t.Errorf("Expected error when trying to load non existent file")
				}
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.test(d)
		})
	}

	// Cleanup file created for testing
	t.Cleanup(func() {
		os.Remove(filename)
	})
}

func TestDeal(t *testing.T) {
	d := newDeck()

	testCases := []struct {
		name string
		test func(d deck)
	}{
		{
			name: "Deal with valid handSize",
			test: func(d deck) {
				hand, remainder := deal(d, 5)
				if len(hand) != 5 {
					t.Errorf("Expected hand with 5 cards, but was %d", len(hand))
				}
				if len(remainder) != 47 {
					t.Errorf("Expected 47 remaining cards, but was %d", len(remainder))
				}
			},
		},
		{
			name: "Deal with invalid handSize",
			test: func(d deck) {
				hand, remainder := deal(d, 55)
				if len(hand) != 52 {
					t.Errorf("Expected hand with 52 cards, but was %d", len(hand))
				}
				if len(remainder) != 0 {
					t.Errorf("Expected 0 remaining cards, but was %d", len(remainder))
				}
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.test(d)
		})
	}
}
