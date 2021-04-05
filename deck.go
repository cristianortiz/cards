package main

import (
	"fmt"
	"strings"
)

//create a new type 'deck' a slice of strings, in this new type we use and expands strings type behavior
type deck []string

//creates a new deck of 52 cards
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	//iterates the slices and make a new deck of cards
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

/* receiver on a function, any variable of type 'deck gets acces to 'print' method
'd' is the receiver of type 'deck' and 'd' is the reference of the deck variable sended to the function */
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//return 2 separates decks of cards
func deal(d deck, handSize int) (deck, deck) {
	//split original d deck, in one deck of cards from index 0 to handSize and other from index handSize to end of the slice
	return d[:handSize], d[handSize:]

}

//reciever to convert a deck type into string
func (d deck) toString() string {
	//d deck to slice of string []string(d)
	//now the slice of string into a single string comma separated ["a","b"]->("a,b")
	return strings.Join([]string(d), ",")

}
