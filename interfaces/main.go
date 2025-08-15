package main

import "fmt"

type bot interface {
	getGreeting(string) string
}
type englishBot struct{}

func main() {
	eb := englishBot{}
	printGreeting(eb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting("Hello"))
}

func (englishBot) getGreeting(s string) string {
	return s
}
