package main

import "fmt"
type deck []string

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

func (cards deck) print(){
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck){
	fmt.Println()

	return d[:handSize], d[handSize:]
}