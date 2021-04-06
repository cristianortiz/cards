package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//create a new type 'deck' a slice of strings, in this new type we use and expands strings type behavior
type deck []string

//creates a new deck of  cards
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	//iterates the slices and make a new deck of cards _index not used in code
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

//receiver to save a deck of cards into a file, te function return an error if thats happends
func (d deck) saveToFile(filename string) error {
	//convert the d from single string and then convert to []byte
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

//reads a plain string file and convert to deck type
func newDeckFromFile(filename string) deck {
	//ReadFile returns a byteSlice from the file, err is nil at least and error is returned
	bs, err := ioutil.ReadFile(filename)
	if err != nil { //error handling in Go
		//1. -log the error and return a call to a newDeck()
		//2. -Log the error and entirely quit the program
		fmt.Println("!Error:", err) //log the error
		os.Exit(1)                  //exit program on error
	}
	//convert the byteSlice, for ex;[50,60,100] into a plain string ("Ace of Spades,Two of Spades")
	s := strings.Split(string(bs), ",") //then split to convert in format ("Ace of Spades","Two of Spades")etc
	//convert the string splited into deck (remember deck is string based type already)
	return deck(s)

}
