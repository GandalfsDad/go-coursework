package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	var fred person
	fred.firstName = "Fred"
	fred.lastName = "Fredson"

	fmt.Printf("%+v", fred)
}
