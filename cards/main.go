package main

func main() {
	//var card string = newCard()
	cards := deck{newCard(), newCard()}

	cards = append(cards, "Six of Spades") //creates new slice

	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
