package main

import (
	"fmt"
)

// Create a new type of 'deck'
// Which is a slice of string

type deck []string

// Create a method 'newDeck'
// Which will return a new 'deck'
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Create a receiver method 'print'
//Which will use 'deck' type varibale to print all the elements inside it

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
