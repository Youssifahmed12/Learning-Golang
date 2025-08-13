package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", newCard()}

	// using append makes a new slice then reassigns it to cards
	cards = append(cards, "King of Spades")

	// using a for loop first "paramater" is the index, second is the value , and use range to iterate over the slice
	for i, card := range cards {
		fmt.Println(card, i)
	}

	// using _ to ignore the index because in golang you can't make a variable without using it
	for _, card := range cards {
		fmt.Println("Card:", card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
