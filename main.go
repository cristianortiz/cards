package main

func main() {
	cards := deck{"Ace of diamonds", newCard()}

	cards = append(cards, "Six of spaces")
	cards.print()
}

func newCard() string {
	return "five of diamonds"
}
