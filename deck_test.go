package main

import "testing"

//t is test handler
func TestNewDeck(t *testing.T) {
	d := newDeck()

	//checks te size of a deck of cards considering the combinatins of slices of suits and values cards
	if len(d) != 16 {
		t.Errorf("Expected deck length of 20, but got", len(d))

	}
}
