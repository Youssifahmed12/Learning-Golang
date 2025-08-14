package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipcode int
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{firstName: "jim",
		lastName: "harbert",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipcode: 1556,
		},
	}
	jim.printInfo()
	jim.updateName("Youssif", "Ahmed")
	jim.printInfo()

}

func (p person) printInfo() {
	fmt.Printf("%+v", p)
}

// using * to pass the struct by refrence so we can actually modify it
func (p *person) updateName(newFirstname string, newLastName string) {
	p.firstName = newFirstname
	p.lastName = newLastName
}
