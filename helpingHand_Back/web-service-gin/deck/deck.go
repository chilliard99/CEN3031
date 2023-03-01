package deck

import (
	//imports card.go as c to prevent redundant card.Card or card.NewCard(int, string) every time a card is created
	c "example/web-service-gin/card"
	//imports hand.go as h to allow hand functions to be called
	h "example/web-service-gin/hand"
	"fmt"
)

// check if string slice already contains item
// from https://freshman.tech/snippets/go/check-if-slice-contains-element/
func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// Define a deck as an array of cards
type Deck []c.Card

func UpdateProb(cards []c.Card, deck Deck) (bool, int) {
	deckCopy := RemoveCards(deck, cards)
	royalBoolean, royalProb := RoyalFlush(deckCopy, cards)

	fmt.Println("Bool result: ", royalBoolean, " Probability: ", royalProb)

	return royalBoolean, int(royalProb)
}

// Determining what hands can be created using the current hand and cards in the deck
func DetermineFutureHands(hand *h.Hand, currentHands []string) []string {
	futureHands := make([]string, 0)
	if len(hand.ActualHand) < 7 {
		if Contains(currentHands, "Full House") {
			futureHands = append(futureHands, "Placeholder - Full House")
		} else if Contains(currentHands, "Three of a Kind") {
			futureHands = append(futureHands, "Four of a Kind")
		} else if Contains(currentHands, "One Pair") {
			futureHands = append(futureHands, "Three of a Kind")
			if len(hand.ActualHand) < 6 {
				futureHands = append(futureHands, "Four of a Kind")
				futureHands = append(futureHands, "Two Pair")
			}
		} else if Contains(currentHands, "Two Pair") {
			futureHands = append(futureHands, "Three of a Kind")
			if len(hand.ActualHand) < 6 {
				futureHands = append(futureHands, "Four of a Kind")
			}
		} else {
			futureHands = append(futureHands, "One Pair")
		}
	} else {
		//no future hands can be made so the function returns "None"
		futureHands = append(futureHands, "None")
	}
	return futureHands
}

// Checking the hand type (i.e. one pair, 3 of a kind...) of the hand at the current time
func CheckHandType(hand *h.Hand) []string {
	doesHandHaveType := false
	handTypes := make([]string, 0) //store hand types
	cardCount := make(map[int]int) //store count of card numbers
	pairCount := 0
	trioCount := 0
	quartetCount := 0
	for _, card := range hand.ActualHand {
		cardCount[card.Val]++
	}
	for i := 0; i < 13; i++ {
		if cardCount[i] == 2 {
			pairCount++
			doesHandHaveType = true
		} else if cardCount[i] == 3 {
			trioCount++
			doesHandHaveType = true
		} else if cardCount[i] == 4 {
			quartetCount++
			doesHandHaveType = true
		}
	}
	hasPair := false
	if !doesHandHaveType {
		handTypes = append(handTypes, "None")
		return handTypes
	}
	if pairCount >= 1 {
		handTypes = append(handTypes, "One Pair")
		hasPair = true
	}
	if pairCount >= 2 {
		handTypes = append(handTypes, "Two Pair")
	}
	if trioCount >= 1 {
		handTypes = append(handTypes, "Three of a Kind")
		if !hasPair {
			handTypes = append(handTypes, "One Pair")
		} else {
			handTypes = append(handTypes, "Full House")
		}
	}
	if quartetCount >= 1 {
		handTypes = append(handTypes, "Four of a Kind")
		handTypes = append(handTypes, "Three of a Kind")
		if !hasPair {
			handTypes = append(handTypes, "One Pair")
		} else {
			handTypes = append(handTypes, "Full House")
		}
	}
	return handTypes
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

func ValSortCardsDes(cards []c.Card) []c.Card {
	cardCount := len(cards)
	swap := true

	for swap {
		swap = false
		for i := 0; i < cardCount-1; i++ {
			if cards[i+1].Val > cards[i].Val {
				tempCard := cards[i+1]
				cards[i+1] = cards[i]
				cards[i] = tempCard
				swap = true
			}
		}
	}

	return cards
}

func ValSortCardsAsc(cards []c.Card) []c.Card {
	cardCount := len(cards)
	swap := true

	for swap {
		swap = false
		for i := 0; i < cardCount-1; i++ {
			if cards[i+1].Val < cards[i].Val {
				tempCard := cards[i+1]
				cards[i+1] = cards[i]
				cards[i] = tempCard
				swap = true
			}
		}
	}

	return cards
}

// Check whether the hand and community cards provided can form a royal straight flush with the rest of the deck and return boolean and probability
func RoyalFlush(deck Deck, cards []c.Card) (bool, float64) {
	deckCopy := RemoveCards(deck, cards)
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

// Returns true if hand + community cards contains a straight and probability. Third boolean represents whether or not the straight is royal
func StraightCheck(deck Deck, cards []c.Card) (bool, float64, bool) {
	//deckCopy := RemoveCards(deck, cards)
	cards = ValSortCardsAsc(cards)
	cardCount := len(cards)
	lowVal := 0
	highVal := 0
	chain := 0

	//VARIABLES FROM ROYALFLUSH (may use later):
	//deckCount := len(deckCopy)
	//remaining := 7 - cardCount
	//lowIndex := 0
	//chosenSuit := ""
	//suits := map[string]struct{}{}
	//straightHand := make(map[int]c.Card)
	//needCards := [5]bool{true, true, true, true, true}
	//var dupe c.Card
	//needCount := 5
	//VARIABLES FROM ROYALFLUSH

	//Checks the cards in hand for a straight by looking for 5 sequentially valued cards
	for i := 0; i < cardCount-1; i++ {
		tempCard := cards[i]
		tempNext := cards[i+1]
		diff := tempNext.Val - tempCard.Val

		//fmt.Println("first: ", tempCard.Val, " second: ", tempNext.Val, " diff: ", diff)

		if diff == 1 {
			if chain == 0 {
				lowVal = tempCard.Val
				highVal = tempNext.Val
				chain = 2

				//fmt.Println("starting chain at: ", lowVal)
			} else {
				chain++
				highVal = tempNext.Val
				//fmt.Println("adding: ", highVal)
			}
		} else {
			chain = 0
			//fmt.Println("restarting chain")
		}
	}

	//ROYAL FLUSH
	if lowVal == 9 && highVal == 12 && chain == 4 && cards[0].Val == 0 {
		return true, 1.00, true //third boolean for royal
	}

	//Regular straight
	if highVal-lowVal == 4 && chain == 5 {
		return true, 1.00, false
	}

	return false, 0.00, false
}

//Started writing flushcheck before focusing on straightcheck
/*
func FlushCheck(deck Deck, cards []c.Card) (bool, int) {
	//deckCopy := RemoveCards(deck, cards)
	cardCount := len(cards)
	suits := make(map[string]int)
	var tempCard c.Card

	//VARIABLES FROM ROYALFLUSH (may use later):
		//deckCount := len(deckCopy)
		//remaining := 7 - cardCount
		//chosenSuit := ""
		//royalFlush := make(map[int]c.Card)
		//needCards := [5]bool{true, true, true, true, true}
		//var dupe c.Card
		//needCount := 5
	//VARIABLES FROM ROYALFLUSH

	//Populate suits map with count of each suit
	for i := 0; i < cardCount; i++ {
		tempCard = cards[i]

		if suits[tempCard.Suit] > 0 {
			suits[tempCard.Suit]++
		} else {
			suits[tempCard.Suit] = 1
		}

		//If there are 3 different suits, then 5/7 cards cannot be the same suit for a flush
		if len(suits) == 3 {
			return false, 0.00
		}
	}
}
*/

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
