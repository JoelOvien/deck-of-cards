package main

import "fmt"

func main() {
	cards := []string{"Five of Diamonds", newCard()}
	cards = append(cards, newCard())
	for i, card := range cards {
		fmt.Println(i, card)
	}
	fmt.Println(cards)
}

func newCard() string {
	return "Ace of Spades"
}
