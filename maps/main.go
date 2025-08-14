package main

import "fmt"

func main() {
	colors := map[string]float64{
		"Youssif": 99.6,
		"Mo'men":  16.6,
		"Mohamed": 13.3,
	}

	printMap(colors)

}

func printMap(c map[string]float64) {
	for n, g := range c {
		fmt.Println("Name :", n, "Grade :", g)
	}
}
