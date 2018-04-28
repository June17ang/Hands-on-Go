package main

func main() {
	cards := newDeckFromFile("my_cards")
	cards.print()

	cards.shuffle()
	cards.print()
}
