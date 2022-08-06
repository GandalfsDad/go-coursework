package main

import "fmt"

func main() {
	//var card string = newCard()
	cards := deck{newCard(), newCard()}

	cards = append(cards, "Six of Spades") //creates new slice

	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func newCard() string {
	return "Five of Diamonds"
}
