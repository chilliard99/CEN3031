package hand

import (
	//imports card.go as c to prevent redundant card.Card or card.NewCard(int, string) every time a card is created
	c "example/web-service-gin/card"
	"math/rand"
	"time"
)

// defining hand as an array of cards separate from the deck
type Hand struct {
	ActualHand []c.Card //the 5 cards in the hand
	HandType   string   //i.e. straight, 4 of a kind, royal flush...
}

func NewHand(handType string) *Hand {
	return &Hand{[]c.Card{}, handType}
}

// add a card to the hand if hand has less than 5 cards
func AddCardHand(hand *Hand) string {
	if len(hand.ActualHand) < 5 {
		//get "random" value from time
		rand.Seed(time.Now().UnixNano())
		Number := rand.Intn(13)
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
		card := c.NewCard(Number, Suit)
		hand.ActualHand = append(hand.ActualHand, card)
		return "successful"
	} else {
		return "length of hand is already 5"
	}
}
