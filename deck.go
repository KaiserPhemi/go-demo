package main

// package import
import (
	"fmt" 
	"strings"
	"io/ioutil"
)

// custom type
type deck []string

// creates a new deck
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Clubs", "Diamonds", "Hearts"}
	cardValues := []string{"Ace","Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _,suit := range cardSuits {
		for _,value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// prints to the standard output
func (cards deck) print(){
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

// splits a deck
func deal(d deck, handSize int) (deck, deck){
	fmt.Println()
	return d[:handSize], d[handSize:]
}

// Converts deck to a single string
func (d deck) toString() string {
		return strings.Join([]string(d), ",")
}

// saves to file
func (d deck) saveToFile(fileName string) error {
	err := ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
	return err
}
