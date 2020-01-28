package main

// import "fmt"

func main() {
	cards := newDeck()
	// cardStrings := cards.toString()
	cards.saveToFile("my_cards")

}
