package main

func main() {
	deck := newDeck()

	hand, remainder := deal(deck, 5)
	hand.print()
	remainder.print()
}

func newCard() string {
	return "Five of Diamonds"
}
