package deck_test

import (
	"example/web-service-gin/deck"
	"testing"
)

// Inputs a card value and suit and returns the index of the card in the deck, comparing it to the value at the end of the if-statement
func TestGetIndex(t *testing.T) {
	//Suits are in the order of Heart (0), Diamond (1), Club (2), Spade (3)
	//This means the index will be the value of the card plus 0, 1, 2, or 3 multiplied by 13 for a full suit
	//5 of Hearts = 5 + (0 * 13) = 5
	//7 of Diamonds = 7 + (1 * 13) = 20

	tempdeck := deck.New()
	if deck.GetCardIndex(tempdeck, 7, "Diamond") != 20 { //index will be the same as the value + (13 * the suit)
		t.Fatal("Wrong card index")
	}
}

// Inputs a card and checks if it returns the appropriate name.
func TestGetCardName(t *testing.T) {
	tempdeck := deck.New()

	//Using GetCardIndex to show the value and suit associated with the selected card
	//val = 12, suit = "Club", card should have name "King of Clubs"
	if deck.GetCardName(tempdeck[deck.GetCardIndex(tempdeck, 12, "Club")]) != "King of Clubs" {
		t.Fatal("Wrong card name")
	}
}
