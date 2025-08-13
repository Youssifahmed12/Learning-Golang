package main

import "fmt"

//Learning how to create Types
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d *deck) addCard() {
	*d = append(*d, "New Card")
}
