package main

func main() {
	cards := newDeck()
	hand, remainingCards := deal(cards, 23)
	hand.print()
	remainingCards.print()
}
