package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	youssif := person{firstName: "youssif", lastName: "ahmed"}
	fmt.Println(youssif)
	youssif.firstName = "Yasser"
	fmt.Println(youssif)

	// prints the struct with the name of the fields
	fmt.Printf("%+v", youssif)
}
