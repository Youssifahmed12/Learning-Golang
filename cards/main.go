package main

import "fmt"

// to make a global variable , you need to use the keyword var not the short :=
var card3 string = "Queen of Diamonds"

/*
	card4 := "Jack of Clubs"
	This will cause an error because := cannot be used at the package level
*/

func main() {
	// Method one to declare a variable
	var card1 string = "Ace of Spades"

	// Method two to declare a variable
	card2 := "King of Hearts"

	fmt.Println(card1)
	fmt.Println(card2)
	fmt.Println(card3)
}
