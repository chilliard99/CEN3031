package deck

import (
	//imports card.go as c to prevent redundant card.Card or card.NewCard(int, string) every time a card is created
	c "example/web-service-gin/card"
	"fmt"
)

// Define a deck as an array of cards
type Deck []c.Card

// Create a new deck filling it with 52 cards: 4 suits, 13 cards each from 0 to 12
func NewDeck() (deck Deck) {

	for i := 0; i < 4; i++ { //4 suits from 0 to 3
		for j := 0; j < 13; j++ { //13 cards from 0 to 12

			//temporary card object with val set as j
			tempCard := c.NewCard(j, "")

			switch i {
			case 0:
				tempCard.Suit = "Heart"
			case 1:
				tempCard.Suit = "Diamond"
			case 2:
				tempCard.Suit = "Club"
			case 3:
				tempCard.Suit = "Spade"
			default:
				tempCard.Suit = "Error in suit assignment" //If you see this appear, something is wrong
			}

			//add the card with the appropriate value and suit to the deck
			deck = append(deck, tempCard)
		}
	}
	return
}

// Function to return the index of a card in the deck using the value and suit
func GetCardIndex(deck Deck, v int, s string) int {
	suitMod := 0

	switch s {
	case "Heart":
		suitMod = 0
	case "Diamond":
		suitMod = 1
	case "Club":
		suitMod = 2
	case "Spade":
		suitMod = 3
	default:
		suitMod = 4 //this will make the next test fail
	}

	//formula for index
	index := v + (suitMod * 13)

	//Probably better ways to throw errors, but setting it to Ace of Hearts seems safe for now
	if index > 51 || v > 12 || v < 0 || suitMod > 3 {
		fmt.Println("Invalid Suit Name")
		index = 0
		return index
	}

	return index
}

// Function to print the entire deck with Index #, card value, and card suit line by line
// ##Currently unable to test with Go test files. See t.Testing log command for a possible solution##
func PrintDeck(deck Deck) {
	for i := 0; i < 52; i++ {
		fmt.Println("Index #", i, ": ", deck[i].Val, " ", deck[i].Suit)
	}
}
