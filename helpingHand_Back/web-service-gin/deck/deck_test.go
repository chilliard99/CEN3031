package deck_test

import (
	"example/web-service-gin/card"
	"example/web-service-gin/deck"
	"strconv"
	"testing"
)

//FOR ALL GetCardIndex TESTS:
// Inputs a card value and suit and returns the index of the card in the deck, comparing it to the value at the end of the if-statement
// Suits are in the order of Heart (0), Diamond (1), Club (2), Spade (3)
// This means the index will be the value of the card plus 0, 1, 2, or 3 multiplied by 13 for a full suit
// 5 of Hearts = 5 + (0 * 13) = 5
// 7 of Diamonds = 7 + (1 * 13) = 20
// King of Clubs = 12 + (2 * 13) = 38
// Ace of Spades = 0 + (3 * 13) = 39

// Tests if input 5, Heart = Five of Hearts (index 5)
func TestGetCardIndex0(t *testing.T) {

	tempdeck := deck.NewDeck()
	t.Log("\n")
	t.Logf("Test #0: GetCardIndex")
	t.Logf("Input of deck, int 5, and string \"Heart\". Comparing to index value 5 (H>D>C>S; suitmod = 0; eq = 5 + (0 * 13))")

	result := deck.GetCardIndex(tempdeck, 5, "Heart")
	t.Logf("result = " + strconv.Itoa(result))
	if result != 5 { //index will be the same as the value + (13 * the suit)
		t.Fatal("Wrong card index")
	}
}

// Tests if input 7, Diamond = Seven of Diamonds (index 20)
func TestGetCardIndex1(t *testing.T) {

	tempdeck := deck.NewDeck()
	t.Log("\n")
	t.Logf("Test #1: GetCardIndex")
	t.Logf("Input of deck, int 7, and string \"Diamond\". Comparing to index value 20 (H>D>C>S; suitmod = 1; eq = 7 + (1 * 13))")

	result := deck.GetCardIndex(tempdeck, 7, "Diamond")
	t.Logf("result = " + strconv.Itoa(result))
	if result != 20 { //index will be the same as the value + (13 * the suit)
		t.Fatal("Wrong card index")
	}
}

// Tests if input 12, Club = King of Clubs (index 38)
func TestGetCardIndex2(t *testing.T) {

	tempdeck := deck.NewDeck()
	t.Log("\n")
	t.Logf("Test #2: GetCardIndex")
	t.Logf("Input of deck, int 12, and string \"Club\". Comparing to index value 38 (H>D>C>S; suitmod = 2; eq = 12 + (2 * 13))")

	result := deck.GetCardIndex(tempdeck, 12, "Club")
	t.Logf("result = " + strconv.Itoa(result))
	if result != 38 { //index will be the same as the value + (13 * the suit)
		t.Fatal("Wrong card index")
	}
}

// Tests if input 0, Spade = Ace of Spades (index 39)
func TestGetCardIndex3(t *testing.T) {

	tempdeck := deck.NewDeck()
	t.Log("\n")
	t.Logf("Test #3: GetCardIndex")
	t.Logf("Input of deck, int 0, and string \"Spade\". Comparing to index value 38 (H>D>C>S; suitmod = 3; eq = 0 + (3 * 13))")

	result := deck.GetCardIndex(tempdeck, 0, "Spade")
	t.Logf("result = " + strconv.Itoa(result))
	if result != 39 { //index will be the same as the value + (13 * the suit)
		t.Fatal("Wrong card index")
	}
}

// Tests input error handling of GetCardIndex function
func TestGetCardIndex4(t *testing.T) {

	tempdeck := deck.NewDeck()
	t.Log("\n")
	t.Logf("Test #4: GetCardIndex error handling")
	t.Logf("Input of deck, int 13, and string \"Foobar\". Comparing to index value 0 (When given an out-of-range value, it should return the first card)")

	result := deck.GetCardIndex(tempdeck, 13, "Foobar")
	t.Logf("result = " + strconv.Itoa(result))
	if result != 0 { //index will be the same as the value + (13 * the suit)
		t.Fatal("Wrong card index")
	}
}

// Inputs a card (via index in the array) and checks if it returns the appropriate name.
func TestGetCardName(t *testing.T) {
	tempdeck := deck.NewDeck()

	//Using GetCardIndex to show the value and suit associated with the selected card
	//val = 12, suit = "Club", card should have name "King of Clubs"
	t.Log("\n")
	t.Logf("Test #5: GetCardName")
	t.Logf("Input of array, index provided by GetCardIndex with input deck, 12, \"Club\". Comparing to card name \"King of Clubs\"")

	result := card.GetCardName(tempdeck[deck.GetCardIndex(tempdeck, 12, "Club")])
	t.Logf("result = " + result)
	if result != "King of Clubs" {
		t.Fatal("Wrong card name")
	}
}

// Prints out the full deck, only visible if running "go test -v"
func TestPrintDeck(t *testing.T) {
	tempdeck := deck.NewDeck()

	t.Log("\n")
	t.Logf("Test #6: PrintDeck")
	t.Logf("Input of deck, prints out full deck in order")

	deck.PrintDeck(tempdeck)
}
