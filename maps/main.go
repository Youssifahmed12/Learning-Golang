package main

import "fmt"

func main() {
	colors := map[string]float64{
		"Youssif": 99.6,
		"Mo'men":  16.6,
	}

	fmt.Println(colors)
	delete(colors, "Mo'men")
	fmt.Println(colors)

}
