package main

import (
	"example/web-service-gin/deck"
	"example/web-service-gin/hand"
	"fmt"
)

func main() {
	currDeck := deck.New()
	fmt.Println(deck.GetCardIndex(currDeck, 0, "Heart"))
	currentHand := hand.NewHand("initial")
	hand.AddCardHand(currentHand)
	fmt.Println("current length of hand is: ")
	fmt.Println(len(currentHand.ActualHand))
}
