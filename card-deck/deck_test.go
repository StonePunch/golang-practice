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