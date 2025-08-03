package main

func main() {
	// cards := newDeck()
	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()

	// create a new deck and save to file
	// cards := newDeck()
	// cards.saveToFile("my_cards")

	// read existing deck from file
	// cards := newDeckFromFile("my_cardss")
	// cards.print()

	// create a new deck, shuffle it and print
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
