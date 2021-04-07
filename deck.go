package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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
	//convert the d deck to single string (toString) and then convert to []byte slice and finally WritesFile create the file
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

//reads a plain string file and convert to deck type
func newDeckFromFile(filename string) deck {
	//ReadFile returns a byteSlice from the file, and err wich is nil at least and error is returned
	bs, err := ioutil.ReadFile(filename)
	if err != nil { //error handling in Go
		//1. -log the error and return a call to a newDeck()
		//2. -Log the error and entirely quit the program
		fmt.Println("!Error:", err) //log the error
		os.Exit(1)                  //exit program on error
	}
	//convert the byteSlice, for ex;[50,60,100] into a plain string ("Ace of Spades,Two of Spades")
	s := strings.Split(string(bs), ",") //then split() to convert in slice []string ("Ace of Spades","Two of Spades")
	//convert the slice []string splited into deck type (rem, deck is a string based type)
	return deck(s)

}

//shuffle receiver function, it will take a deck of cards and randomized it
func (d deck) shuffle() {
	/* 	1. create a trully random generator, starts whit a source object (is like a seed)
	   	every sigle time the code runs use the Time object to create a unique value as a source (or seed)
	   	for random number generator, NewSource needs and int64 as param, and UnixNano returns a int64 from actual time */
	source := rand.NewSource(time.Now().UnixNano())
	//using the source object creates a random generator object *Rand type
	r := rand.New(source)
	//2. iterate the deck of cards, we will use only the index of the slice
	for i := range d {
		//*Rand Object has Intn() method to creates a trully a random int beetwen 0 and lenght of slice - 1
		newPosition := r.Intn(len(d) - 1)
		//4. in every i slice position, swap the values of the indexes, newPosition is a random number
		//so, the position of cards always will be random
		d[i], d[newPosition] = d[newPosition], d[i]

	}

}

/*Remember reciever functions allow us to use data.fun() like a method of an object notation and
aply the funcion directly to the object*/
