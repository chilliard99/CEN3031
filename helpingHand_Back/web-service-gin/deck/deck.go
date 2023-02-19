package deck

import (
	//imports card.go as c to prevent redundant card.Card or card.NewCard(int, string) every time a card is created
	c "example/web-service-gin/card"
	"fmt"
)

// Define a deck as an array of cards
type Deck []c.Card

func UpdateProb(cards []c.Card, deck Deck) {

}

// Removes hand and community cards from the deck. Returns deckCopy without the input cards
func RemoveCards(deck Deck, cards []c.Card) (deckCopy Deck) {
	deckCount := len(deck)
	cardCount := len(cards)

	//Check each card of the deck against input. Do not add to deckCopy if card exists in hand or community cards.
	for j := 0; j < deckCount; j++ {
		addCheck := true

		for i := 0; i < cardCount; i++ {
			if cards[i].Val == deck[j].Val && cards[i].Suit == deck[j].Suit {
				addCheck = false
			}
		}

		if addCheck {
			deckCopy = append(deckCopy, deck[j])
		}
	}

	return
}

// Check whether the hand and community cards provided can form a royal straight flush with the rest of the deck and return boolean and probability
func RoyalFlush(deckCopy Deck, cards []c.Card) (bool, float64) {
	deckCount := len(deckCopy)
	cardCount := len(cards)
	remaining := 7 - cardCount

	/* POSSIBLY NOT NEEDED
	turnCount := 0

	//determine which turn the game is in based on hand cards + community cards
	switch cardCount {
	case 0:
		turnCount = 0 //Buy-in
	case 2:
		turnCount = 1 //Opening bet (after player receives 2 cards)
	case 5:
		turnCount = 2 //Flop (first 3 community cards flipped)
	case 6:
		turnCount = 3 //Turn (4th community card)
	case 7:
		turnCount = 4 //River (2 cards in hand + 5 community cards)
	default:
		turnCount = 0 //Should never be assigned unless entering cards outside of game rules
	}
	*/

	chosenSuit := ""
	suits := map[string]struct{}{}
	royalFlush := make(map[int]c.Card)
	needCards := [5]bool{true, true, true, true, true}

	var dupe c.Card
	var tempCard c.Card

	needCount := 5

	//Populate royalFlush map with cards of like suit for values of 12, 11, 10, 9, 0
	for i := 0; i < cardCount; i++ {
		tempCard = cards[i]

		//***THIS PART COULD BE MADE INTO A FLUSHCHECK FUNCTION***
		suits[tempCard.Suit] = struct{}{}

		//If there are 3 different suits, then 5/7 cards cannot be the same suit for a flush
		if len(suits) == 3 {
			return false, 0.00
		}
		//***THIS PART COULD BE MADE INTO A FLUSHCHECK FUNCTION***

		//Switch case to ignore all cards not needed for a royal flush straight
		switch tempCard.Val {
		case 12, 11, 10, 9, 0:

			//If map slot is empty, add card
			if royalFlush[tempCard.Val].Suit == "" {
				royalFlush[tempCard.Val] = tempCard
				chosenSuit = tempCard.Suit

				//Check dupe to make sure cards are in the right suit (in case of duplicate value card in the wrong suit)
				if dupe.Suit == tempCard.Suit {
					royalFlush[dupe.Val] = dupe
					dupe.Suit = ""
				}

				needCount--

				//mark card off on bool array
				switch tempCard.Val {
				case 12:
					needCards[0] = false
				case 11:
					needCards[1] = false
				case 10:
					needCards[2] = false
				case 9:
					needCards[3] = false
				case 0:
					needCards[4] = false

				}

				//If map slot is not empty
			} else {
				dupe = tempCard
			}
		default:
		}
	}

	//if suits are not all the same, return false and 0.0%
	for i := 0; i < 5; i++ {

		//check if card is found from the 5 needed
		if !needCards[i] {
			for j := 0; j < 5; j++ {

				//check if card is found from other 4 needed
				if j != i && !needCards[j] {

					//set indexes
					firstIndex := 0
					secondIndex := 0

					switch i {
					case 0:
						firstIndex = 12
					case 1:
						firstIndex = 11
					case 2:
						firstIndex = 10
					case 3:
						firstIndex = 9
					case 4:
						firstIndex = 0
					}

					switch j {
					case 0:
						firstIndex = 12
					case 1:
						firstIndex = 11
					case 2:
						firstIndex = 10
					case 3:
						firstIndex = 9
					case 4:
						firstIndex = 0
					}

					//compare card suits
					if royalFlush[firstIndex].Suit != royalFlush[secondIndex].Suit {
						return false, 0.00
					}
				}
			}
		}
	}

	//If more cards are needed than should be drawn, return false and 0.0%
	if needCount > remaining {
		return false, 0.00
	} else if needCount == 0 {
		return true, 1.00 //if none needed, return true and 100.0%
	}

	totalProb := 0.00
	//Calculate possibilities of getting other cards from the deck
	for remaining > 0 {
		boolIndex := 0
		firstVal := 99
		tempProb := 0.00

		//Get the proper card value from boolean array
		for i := 0; i < 5; i++ {
			if needCards[i] {
				switch i {
				case 0:
					firstVal = 12
					boolIndex = 0
				case 1:
					firstVal = 11
					boolIndex = 1
				case 2:
					firstVal = 10
					boolIndex = 2
				case 3:
					firstVal = 9
					boolIndex = 3
				case 4:
					firstVal = 0
					boolIndex = 4
				default:
				}
				break
			}
		}

		//Go through deck, find the needed card, remove it, and add to the probability
		for i := 0; i < deckCount; i++ {
			if deckCopy[i].Suit == chosenSuit && deckCopy[i].Val == firstVal {
				//Find probability
				tempProb = 1.00 / float64(deckCount)

				//Adjust boolean array and other variables for found card
				needCards[boolIndex] = false
				remaining--

				//Remove the card from the deckCopy and break
				var rem []c.Card
				rem = append(rem, deckCopy[i])
				deckCopy = RemoveCards(deckCopy, rem)
				break
			}

		}

		//needed card not found in deck, return false and 0.0%
		if tempProb == 0.00 {
			return false, 0.00
		}

		//Set totalProb if not already set
		if totalProb == 0.00 {
			totalProb = tempProb
			tempProb = 0.00

			//THIS CALCULATION ASSUMES DRAWING ALL CARDS IN THE ORDER OF CALCULATION (for 2 cards in any order, it is multiplied by 2. For 3 in any order, it is multiplied by 6, etc.)
		} else {
			totalProb = totalProb * tempProb
		}
	}

	//I have no idea if this accounts for the number of permutations/orderings for cards
	switch needCount {
	case 1:
	case 2:
		totalProb = totalProb * 2
	case 3:
		totalProb = totalProb * 6
	case 4:
		totalProb = totalProb * 24
	case 5:
		totalProb = totalProb * 120
	case 6:
		totalProb = totalProb * 720
	case 7:
		totalProb = totalProb * 5040
	}

	return true, totalProb
}

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
		fmt.Println("Invalid range (int 0-12, suit Heart/Diamond/Club/Spade)")
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
