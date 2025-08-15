package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}

func main() {
	eb := englishBot{}
	printGreeting(eb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hello"
}
