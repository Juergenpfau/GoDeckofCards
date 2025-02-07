package main

import (
	"fmt"
)

func main() {
	cards := deck{"Ace Of Diamonds", newCard()}
	fmt.Println(cards)
	cards = append(cards, "Six Of Spades")

	for index, card := range cards {
		fmt.Println(index, card)
	}
}

func newCard() string {
	return "Five Of Diamonds"
}
