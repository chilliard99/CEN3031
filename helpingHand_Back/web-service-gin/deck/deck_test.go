package deck_test

import (
	"example/web-service-gin/card"
	"example/web-service-gin/deck"
	"example/web-service-gin/hand"
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

// Removes 3 cards from a copy of the deck, then tests the removal by running GetCardIndex on a full deck to see if the indexes were offset
func TestRemoveCards(t *testing.T) {
	tempDeck := deck.NewDeck()

	card1 := card.NewCard(0, "Heart")
	card2 := card.NewCard(12, "Club")
	card3 := card.NewCard(2, "Spade")

	var cards []card.Card

	cards = append(cards, card1)
	cards = append(cards, card2)
	cards = append(cards, card3)

	t.Log("\n")
	t.Logf("Test #7: RemoveCards")
	t.Logf("Input of deck, selection of 3 cards, comparing to full name of card with appropriate index offset")

	deckCopy := deck.RemoveCards(tempDeck, cards)

	if card.GetCardName(deckCopy[0]) != "Two of Hearts" { //value should be 1 higher (1 card removed by this index) Ace of Hearts -> Two of Hearts
		t.Fatal("Card 1 not removed")
	}
	if card.GetCardName(deckCopy[deck.GetCardIndex(tempDeck, 12, "Club")]) != "Two of Spades" { //value should be 2 higher (2 cards removed by this index) King of Clubs -> Ace of Spades -> Two of Spades
		t.Fatal("Card 2 not removed")
	}
	if card.GetCardName(deckCopy[deck.GetCardIndex(tempDeck, 2, "Spade")]) != "Six of Spades" { //value should be 3 higher (3 cards removed by this index) Three of Spades -> 4ofS -> 5ofS -> 6ofS
		t.Fatal("Card 3 not removed")
	}
}

// **ONLY TESTS FOR ROYALFLUSH PRESENCE CURRENTLY**
// Test to check whether the RoyalFlush function can return a true output given 4 of 5 cards required.
func TestRoyalFlushCheck(t *testing.T) {
	tempDeck := deck.NewDeck()

	card1 := card.NewCard(12, "Spade")
	card2 := card.NewCard(11, "Spade")
	card3 := card.NewCard(10, "Spade")
	card4 := card.NewCard(9, "Spade")
	card5 := card.NewCard(0, "Spade") //added line for half-measure test

	var cards []card.Card

	cards = append(cards, card1)
	cards = append(cards, card2)
	cards = append(cards, card3)
	cards = append(cards, card4)
	cards = append(cards, card5) //added line for half-measure test

	t.Log("\n")
	t.Logf("Test #8: RoyalFlushCheck")
	//t.Logf("Input of deck, selection of 4 cards (for royal flush), output should be true as it only need ace of spades")

	t.Logf("Input of deck, selection of 5 cards (for royal flush), output should be true") //added line for half-measure test

	boolResponse, probFloat := deck.RoyalFlush(tempDeck, cards)

	if boolResponse == false {
		t.Fatal("Returned false when it should've returned true")
	}
	if probFloat == 0.00 {
		t.Fatal("Returned 0.0 when it should be greater")
	}
}

// Test to check whether straight function will properly identify a straight (royal)
func TestStraightCheck1(t *testing.T) {
	tempDeck := deck.NewDeck()

	card1 := card.NewCard(12, "Spade")
	card2 := card.NewCard(11, "Spade")
	card3 := card.NewCard(10, "Spade")
	card4 := card.NewCard(9, "Spade")
	card5 := card.NewCard(0, "Spade")

	var cards []card.Card

	cards = append(cards, card1)
	cards = append(cards, card2)
	cards = append(cards, card3)
	cards = append(cards, card4)
	cards = append(cards, card5)

	t.Log("\n")
	t.Logf("Test #9: StraightCheck (royal)")
	t.Logf("Input of deck, selection of 5 cards (for royal flush), output should be true")

	boolResponse, probFloat, royalBoolean := deck.StraightCheck(tempDeck, cards)

	if boolResponse == false {
		t.Fatal("Returned false when it should've returned true")
	}
	if probFloat == 0.00 {
		t.Fatal("Returned 0.0 when it should be greater")
	}
	if royalBoolean == false {
		t.Fatal("Returned false for royal when it should've returned true")
	}
}

// Test to check whether straight function will properly identify a straight (non-royal)
func TestStraightCheck2(t *testing.T) {
	tempDeck := deck.NewDeck()

	card1 := card.NewCard(11, "Spade")
	card2 := card.NewCard(10, "Spade")
	card3 := card.NewCard(9, "Spade")
	card4 := card.NewCard(8, "Spade")
	card5 := card.NewCard(7, "Spade")

	var cards []card.Card

	cards = append(cards, card1)
	cards = append(cards, card2)
	cards = append(cards, card3)
	cards = append(cards, card4)
	cards = append(cards, card5)

	t.Log("\n")
	t.Logf("Test #10: StraightCheck (non-royal)")
	t.Logf("Input of deck, selection of 5 sequential cards, output should be true")

	boolResponse, probFloat, royalBoolean := deck.StraightCheck(tempDeck, cards)

	if boolResponse == false {
		t.Fatal("Returned false when it should've returned true")
	}
	if probFloat == 0.00 {
		t.Fatal("Returned 0.0 when it should be greater")
	}
	if royalBoolean == true {
		t.Fatal("Returned true for royal when it should've returned false")
	}
}

// Test to check accuracy of card array sorting functions
func TestValSortCards(t *testing.T) {
	card1 := card.NewCard(1, "Heart")
	card2 := card.NewCard(11, "Spade")
	card3 := card.NewCard(10, "Club")
	card4 := card.NewCard(4, "Diamond")
	card5 := card.NewCard(8, "Spade")

	var cards []card.Card

	cards = append(cards, card1)
	cards = append(cards, card2)
	cards = append(cards, card3)
	cards = append(cards, card4)
	cards = append(cards, card5)

	sortedAsc := deck.ValSortCardsAsc(cards)

	lastVal := 0

	for i := 0; i < len(cards); i++ {
		t.Logf(strconv.Itoa(sortedAsc[i].Val))
		if lastVal > sortedAsc[i].Val {
			t.Fatal("Cards are not in ascending order")
		}

		lastVal = sortedAsc[i].Val
	}

	sortedDes := deck.ValSortCardsDes(cards)

	lastVal = 13

	for i := 0; i < len(cards); i++ {
		t.Logf(strconv.Itoa(sortedDes[i].Val))
		if lastVal < sortedDes[i].Val {
			t.Fatal("Cards are not in descending order")
		}

		lastVal = sortedDes[i].Val
	}
}

// Attempts to see if hand type changes after inserting one pair of cards
func TestOnePairCheck(t *testing.T) {
	t.Log("Testing one pair functionality")
	temphand := hand.NewHand("None")
	hand.AddCardHandSpecific(temphand, 1, "Heart")
	hand.AddCardHandSpecific(temphand, 2, "Heart")
	hand.AddCardHandSpecific(temphand, 1, "Club")
	hand.AddCardHandSpecific(temphand, 3, "Spade")
	hand.AddCardHandSpecific(temphand, 4, "Diamond")
	if deck.Contains(deck.CheckHandType(temphand), "One Pair") != true {
		t.Fatal("One Pair comparison is not working!")
	} else {
		t.Log("One Pair comparison successful!")
	}
}

// Attempts to see if hand type changes after inserting two pairs of cards
func TestTwoPairCheck(t *testing.T) {
	t.Log("Testing two pair functionality")
	temphand := hand.NewHand("None")
	hand.AddCardHandSpecific(temphand, 1, "Heart")
	hand.AddCardHandSpecific(temphand, 2, "Heart")
	hand.AddCardHandSpecific(temphand, 1, "Club")
	hand.AddCardHandSpecific(temphand, 2, "Club")
	hand.AddCardHandSpecific(temphand, 4, "Diamond")
	if deck.Contains(deck.CheckHandType(temphand), "Two Pair") != true {
		t.Fatal("Two Pair comparison is not working!")
	} else {
		t.Log("Two Pair comparison successful!")
	}
}

// Attempts to see if hand type changes after inserting three and four of a kind of cards
// Also full house now
func TestThreeFourFullCheck(t *testing.T) {
	t.Log("Testing three of a kind functionality")
	temphand := hand.NewHand("None")
	hand.AddCardHandSpecific(temphand, 1, "Heart")
	hand.AddCardHandSpecific(temphand, 1, "Spade")
	hand.AddCardHandSpecific(temphand, 1, "Club")
	hand.AddCardHandSpecific(temphand, 4, "Diamond")
	if deck.Contains(deck.CheckHandType(temphand), "Three of a Kind") != true {
		t.Fatal("Three of a Kind comparison is not working!")
	} else {
		t.Log("Three of a Kind comparison successful!")
	}
	hand.AddCardHandSpecific(temphand, 1, "Diamond")
	if deck.Contains(deck.CheckHandType(temphand), "Four of a Kind") != true {
		t.Fatal("Four of a Kind comparison is not working!")
	} else {
		t.Log("Four of a Kind comparison successful!")
	}
	hand.AddCardHandSpecific(temphand, 4, "Spade")
	if deck.Contains(deck.CheckHandType(temphand), "Full House") != true {
		t.Fatal("Full House comparison is not working!")
	} else {
		t.Log("Full House comparison successful!")
	}
}

// Trying to see if future hand function works
func TestFutureHand(t *testing.T) {
	t.Log("Testing future hand functionality")
	temphand := hand.NewHand("None")
	hand.AddCardHandSpecific(temphand, 1, "Heart")
	hand.AddCardHandSpecific(temphand, 1, "Spade")
	hand.AddCardHandSpecific(temphand, 1, "Club")
	if deck.Contains(deck.DetermineFutureHands(temphand, deck.CheckHandType(temphand)), "Four of a Kind") != true {
		t.Fatal("Future hand function does not work for three of a kind!")
	}
	t.Log(("Future hand function works for three of a kind!"))
}
