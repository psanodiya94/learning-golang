package main

func main() {
	cards := newDeck()
	cards.print()

	cards.saveToFile("my_cards")

	cards2 := newDeckFromFile("my_cards")
	cards2.shuffle()
	cards2.print()

	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()
}
