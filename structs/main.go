package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	fred := person{
		firstName: "Fred",
		lastName:  "FredSon",
		contact: contactInfo{
			email:   "fred@fred.fred",
			zipCode: 69420,
		},
	}

	fmt.Printf("%+v", fred)
}
