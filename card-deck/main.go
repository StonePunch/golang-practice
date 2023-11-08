package main

func main() {
	deck := newDeck()

	hand, remainder := deal(deck, 5)
	hand.print()
	remainder.print()

	_ = deck.saveToFile("my_deck")
	loadedDeck, _ := newDeckFromFile("my_deck")

	loadedDeck.shuffle()
	loadedDeck.print()
}
