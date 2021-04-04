package main

import "fmt"

//create a new type 'deck' a slice of strings, in this new type we use and expands strings type behavior
type deck []string

/* receiver on a function, any variable of type 'deck gets acces to 'print' method
'd' is the receiver of type 'deck' and 'd' is the reference of the deck variable sended to the function */
func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}

}
