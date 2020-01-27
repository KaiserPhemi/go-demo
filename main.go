package main

import "fmt"

func main() {
	cards := newDeck()
	cardStrings := cards.toString()
	fmt.Println(cardStrings)

}
