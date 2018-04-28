package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

// Create a method 'deal'
//That will simulate the dealing of a deck of cards with a specific handsize

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//Create a receiver method 'toString'
//Which takes a 'deck' and returns a string

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

//Create a receiver method 'saveToFile'
//Which will save the deck into the hard drive as a file
//This method takes an argument 'filename' of type string

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

//Create a method 'newDeckFromFile'
//Which will read from a file and return a 'deck'
//This method takes an argument 'filename' of type string

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}

//Create a receiver method 'shuffle'
//Which will random shuffle the cards of a 'deck'

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
