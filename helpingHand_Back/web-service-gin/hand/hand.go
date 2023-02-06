package hand

import (
	"fmt"
	"math/rand"
	"time"
)

// copied Card structure from deck.go for consistency
type Card struct {
	Val  int    `json:"val"`
	Suit string `json:"suit"`
}

// defining hand as an array of cards separate from the deck
type Hand struct {
	ActualHand []Card //the 5 cards in the hand
	HandType   string //i.e. straight, 4 of a kind, royal flush...
}

func NewHand(handType string) *Hand {
	return &Hand{[]Card{}, handType}
}

// add a card to the hand if hand has less than 5 cards
func AddCardHand(hand *Hand) {
	if len(hand.ActualHand) < 5 {
		//get "random" value from time
		rand.Seed(time.Now().UnixNano())
		Number := rand.Intn(14)
		SuitValue := rand.Intn(4)
		Suit := ""
		switch SuitValue {
		case 0:
			Suit = "Heart"
		case 1:
			Suit = "Diamond"
		case 2:
			Suit = "Club"
		case 3:
			Suit = "Spade"
		default:
			Suit = "Error when adding card to hand"
		}
		card := Card{Number, Suit}
		hand.ActualHand = append(hand.ActualHand, card)
	} else {
		fmt.Println("length of hand is already 5")
	}
}
