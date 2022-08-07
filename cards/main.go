package main

func main() {
	cards := newDeck()

	hand, remainder := deal(cards, 5)

	hand.print()
	remainder.print()
}
