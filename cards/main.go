package main

func main() {
	cards := newDeck()

	cards.shuffle()

	hand, remainingCards := deal(cards, 5)

	hand.saveToFile("hand")
	remainingCards.saveToFile("remaining")

	oldHand := newDeckFromFile("hand")
	oldHand.print()
}
