package main

import (
	"os"
	"testing"
)

//t is test handler
func TestNewDeck(t *testing.T) {
	d := newDeck()

	//checks te size of a deck of cards considering the combinatins of slices of suits and values cards
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))

	}
	//test the right order of cardSuits elements
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of 'Ace of Spades', but got %v", d[0])
	}

	//the last element in the deck of cards must be a four of clubs
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of 'Four of Clubs', but got %v", d[len(d)-10])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	//delete any previous testing file
	os.Remove("_decktesting")

	//create a new deck and save it in disk
	d := newDeck()
	d.saveToFile("_decktesting")

	//now load the file to create a new deck
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got%v ", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
