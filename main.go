package main

func main() {
	//creates a new decks of cards
	//cards := newDeck()
	//split the deck of cards in two hands of fixed size
	//hand, remainingCards := deal(cards, 5)
	//print the cards in both hands
	//cards.saveToFile("my_deck")
	//read a file and create a new deck of cards from it catches read errors too
	//cards := newDeckFromFile("my_deck")
	cards := newDeck()
	cards.shuffle()
	cards.print()

}
