package main

import "fmt"

func main() {
	//var card string = newCard()
	cards := []string{newCard(), newCard()}

	cards = append(cards, "Six of Spades") //creates new slice

	for i, card := range cards {
		fmt.Println(i, card)
	}
	//fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
