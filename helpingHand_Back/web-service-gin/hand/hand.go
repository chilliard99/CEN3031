package hand

import (
	"math/rand"
	"strconv"
	"time"
)

// copied Card structure from deck.go for consistency
type Card struct {
	Val  string `json:"val"`
	Suit string `json:"suit"`
}

// defining hand as an array of cards separate from the deck
type Hand struct {
	ActualHand []Card //the 5 cards in the hand
	HandType   string //i.e. straight, 4 of a kind, royal flush...
}

type Getter interface {
	GetAll() []Card
}

type Adder interface {
	Add(card Card)
}

// array of cards inputted by user
type UserHand struct {
	Cards []Card
}

func New() *UserHand {
	return &UserHand{
		Cards: []Card{},
	}
}

func (r *UserHand) Add(card Card) {
	r.Cards = append(r.Cards, card)
}

func (r *UserHand) GetAll() []Card {
	return r.Cards
}

func NewHand(handType string) *Hand {
	return &Hand{[]Card{}, handType}
}

// add a card to the hand if hand has less than 5 cards
func AddCardHand(hand *Hand) string {
	if len(hand.ActualHand) < 5 {
		//get "random" value from time
		rand.Seed(time.Now().UnixNano())
		NumberValue := rand.Intn(14)
		Number := ""
		SuitValue := rand.Intn(4)
		Suit := ""
		switch NumberValue {
		case 0:
			Number = "Ace"
		case 11:
			Number = "Jack"
		case 12:
			Number = "Queen"
		case 13:
			Number = "King"
		default:
			Number = strconv.Itoa(NumberValue + 1)
		}
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
		return "successful"
	} else {
		return "length of hand is already 5"
	}
}
