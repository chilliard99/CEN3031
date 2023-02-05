package deck

import "fmt"

//Defining a Card structure with value and suit
type Card struct {
	val  int
	suit string
}

//Define a deck as an array of cards
type Deck []Card

//Uses value and suit to return the name of the Card (i.e. "King of Clubs")
func GetCardName(card Card) string {
	value := ""

	//Simple switch case to translate 0-12 to proper card names
	switch card.val {
	case 0:
		value = "Ace"
	case 1:
		value = "Two"
	case 2:
		value = "Three"
	case 3:
		value = "Four"
	case 4:
		value = "Five"
	case 5:
		value = "Six"
	case 6:
		value = "Seven"
	case 7:
		value = "Eight"
	case 8:
		value = "Nine"
	case 9:
		value = "Ten"
	case 10:
		value = "Jack"
	case 11:
		value = "Queen"
	case 12:
		value = "King"
	default:
		value = "error"
	}

	return value + " of " + card.suit + "s"
}

//Create a new deck filling it with 52 cards: 4 suits, 13 cards each from 0 to 12
func New() (deck Deck) {

	for i := 0; i < 4; i++ { //4 suits from 0 to 3
		for j := 0; j < 13; j++ { //13 cards from 0 to 12

			//temporary card object with val set as j
			tempCard := Card{j, ""}

			switch i {
			case 0:
				tempCard.suit = "Heart"
			case 1:
				tempCard.suit = "Diamond"
			case 2:
				tempCard.suit = "Club"
			case 3:
				tempCard.suit = "Spade"
			default:
				tempCard.suit = "Joker" //If you see this appear, something is wrong
			}

			//add the card with the appropriate value and suit to the deck
			deck = append(deck, tempCard)
		}
	}
	return
}

//Function to return the index of a card in the deck using the value and suit
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
		suitMod = 4
	}

	//formula for index
	index := v + (suitMod * 13)

	//Probably better ways to throw errors, but setting it to Ace of Hearts seems safe for now
	if index > 51 {
		fmt.Println("Invalid Suit Name")
		index = 0
		return index
	}

	return index
}

//Function to print the entire deck with Index #, card value, and card suit line by line
//##Currently unable to test with Go test files. See t.Testing log command for a possible solution##
func PrintDeck(deck Deck) {
	for i := 0; i < 52; i++ {
		fmt.Println("Index #", i, ": ", deck[i].val, " ", deck[i].suit)
	}
}
