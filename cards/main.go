package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades" //same as above but shortcut

	cards := newDeck()

	cards.saveToFile("mycards")

	fmt.Println(cards.toString())

	hand, remainingDeck := deal(cards, 5)

	cardsFromFile := newDeckFromFile("mycards")
	cardsFromFile.print()
	hand.print()
	remainingDeck.print()

	cards.shuffle()
	cards.print()
}
