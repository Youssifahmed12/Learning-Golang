package main

import "fmt"

func main() {
	card := newDeck()
	card.shuffle()
	card.print()
	fmt.Println("==============")
	card.shuffle()
	card.print()
}
