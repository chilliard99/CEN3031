package main

import (
	"example/web-service-gin/deck"
	"fmt"
)

func main() {
	currDeck := deck.New()
	fmt.Println(deck.GetCardIndex(currDeck, 0, "Heart"))
}
