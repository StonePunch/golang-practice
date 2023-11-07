package main

func main() {
	deck := newDeck()

	hand, remainer := deal(deck, 5)
	hand.print()
	remainer.print()
}

func newCard() string {
	return "Five of Diamonds"
}
