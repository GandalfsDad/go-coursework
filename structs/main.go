package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	fred := person{
		firstName: "Fred",
		lastName:  "FredSon",
		contactInfo: contactInfo{
			email:   "fred@fred.fred",
			zipCode: 69420,
		},
	}

	fred.updateName("Bill")
	fred.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
