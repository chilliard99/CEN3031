package hand

import (
	c "example/web-service-gin/card"
	//"strconv"
	"math/rand"
	"time"
)

// Defining hand as an array of cards separate from the deck
type Hand struct {
	ActualHand []c.Card //the 5 cards in the hand
	HandType   string   //i.e. straight, 4 of a kind, royal flush...
}

type Getter interface {
	GetAll() []c.Card
}

type Adder interface {
	Add(card c.Card)
}

// array of cards inputted by user
type UserHand struct {
	//index int
	Cards    []c.Card
	HandType string
}

func New() *UserHand {
	return &UserHand{
		Cards: []c.Card{},
	}
}

func (r *UserHand) Add(card c.Card) {
	temp := r.GetAll()
	//check validity of input
	for i := 0; i < len(temp); i++ {
		if temp[i].Suit == card.Suit && temp[i].Val == card.Val {
			return
		}
	}
	for i := 0; i < len(temp); i++ {
		if temp[i].Index == card.Index {
			temp[i] = card
			return
		}
	}
	r.Cards = append(r.Cards, card)
}

/*
func (r *UserHand) Add(card c.Card, index int) {
	r.Cards[index] = card
}
*/

func (r *UserHand) GetAll() []c.Card {
	return r.Cards
}

// Creating a new hand
func NewHand(handType string) *Hand {
	return &Hand{[]c.Card{}, handType}
}

// Checking the card at an index in the hand
func CheckCardIndex(hand *Hand, index int) c.Card {
	return hand.ActualHand[index]
}

// Add a specific card to the hand if hand has less than 8 cards ,for testing only
func AddCardHandSpecific(hand *Hand, val int, suit string) string {
	if val < 0 || val > 12 {
		return "adding specific card: value is invalid"
	} else if suit != "Heart" && suit != "Diamond" && suit != "Club" && suit != "Spade" {
		return "adding specific card: suit is invalid"
	} else {
		if len(hand.ActualHand) < 8 {
			newCard := c.Card{
				Val:  val,
				Suit: suit,
			}
			hand.ActualHand = append(hand.ActualHand, newCard)
			return "adding specific card: successful"
		} else {
			return "adding specific card: length of hand is already 7"
		}
	}
}

// Add a random card to the hand if hand has less than 5 cards
func AddCardHandRandom(hand *Hand) string {
	if len(hand.ActualHand) < 7 {
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
			Suit = "Error when adding card to hand" //this should not ever be the value so if it is then an error has occured
		}
		//reverted to old code to get merge to work
		//card := c.NewCard(Number, Suit)
		newCard := c.Card{Val: Number, Suit: Suit, Index: Number}
		hand.ActualHand = append(hand.ActualHand, newCard)
		return "adding random card: successful"
	} else {
		return "adding random card: length of hand is already 7"
	}
}
