package main

import "fmt"

func main() {
	//creates a new decks of cards
	cards := newDeck()
	//split the deck of cards in two hands of fixed size
	//hand, remainingCards := deal(cards, 5)
	//print the cards in both hands
	fmt.Println(cards.toString())

}
