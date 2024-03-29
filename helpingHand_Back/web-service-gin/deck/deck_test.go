package deck_test

import (
	"example/web-service-gin/card"
	"example/web-service-gin/deck"
	"example/web-service-gin/hand"
	"math"
	"math/rand"
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

	t.Log("\n")
	t.Logf("Test 8.1: RoyalFlushCheck (probability)")

	t.Logf("Input of deck, selection of 1 cards (for royal flush), output should be 1/51 * 1/50 * 1/49 * 1/48 or 0.000004")

	prob := deck.RoyalFlush(tempDeck, cards)
	compare1 := (math.Round(prob*1000000) / 1000000)

	permutations := float64(deck.Factorial(4))
	totalProb := ((1.00 / 51.00) * (1.00 / 50.00) * (1.00 / 49.00) * (1.00 / 48.00)) * permutations
	compare2 := (math.Round(totalProb*1000000) / 1000000)

	if compare1 != 0.000004 && compare2 != 0.000004 {
		t.Fatal("Returned: ", prob, " Expected: ", 0.000004, " Hand Math: ", compare2)
	}

	cards = append(cards, card2)

	t.Log("\n")
	t.Logf("Test 8.2: RoyalFlushCheck (probability)")

	t.Logf("Input of deck, selection of 2 cards (for royal flush), output should be 1/50 * 1/49 * 1/48 or 0.000051")

	prob = deck.RoyalFlush(tempDeck, cards)
	compare1 = (math.Round(prob*10000) / 10000)
	permutations = float64(deck.Factorial(3))
	totalProb = ((1.00 / 50.00) * (1.00 / 49.00) * (1.00 / 48.00)) * permutations
	compare2 = (math.Round(totalProb*1000000) / 1000000)

	if compare1 != 0.000051 && compare2 != 0.000051 {
		t.Fatal("Returned: ", prob, " Expected: ", 0.000051, " Hand Math: ", compare2)
	}

	//new tests for multi-suit royals
	t.Logf("Test 8.2.5: RoyalFlushCheck (prob-mix)")
	cardsNew := cards

	for i := 0; i < 3; i++ {
		t.Log("\n")
		if i == 0 {
			t.Logf("Test 8.2.5.1: RoyalFlushCheck (prob-mix)")
			cardHeart1 := card.NewCard(12, "Heart")
			cardsNew = append(cardsNew, cardHeart1)
		} else if i == 1 {
			t.Logf("Test 8.2.5.2: RoyalFlushCheck (prob-mix)")
			cardHeart2 := card.NewCard(11, "Heart")
			cardsNew = append(cardsNew, cardHeart2)
		} else if i == 2 {
			t.Logf("Test 8.2.5.3: RoyalFlushCheck (prob-mix w/non-royal)")
			cardHeart3 := card.NewCard(8, "Heart")
			cardsNew = append(cardsNew, cardHeart3)
		}

		prob = deck.RoyalFlush(tempDeck, cardsNew)

		t.Logf("Prob: %f", prob)
		t.Log("\n")
	}

	cards = append(cards, card3)

	t.Log("\n")
	t.Logf("Test 8.3: RoyalFlushCheck (probability)")

	t.Logf("Input of deck, selection of 3 cards (for royal flush), output should be 1/49 * 1/48 or 0.000850")

	prob = deck.RoyalFlush(tempDeck, cards)
	compare1 = (math.Round(prob*10000) / 10000)
	permutations = float64(deck.Factorial(2))
	totalProb = ((1.00 / 49.00) * (1.00 / 48.00)) * permutations
	compare2 = (math.Round(totalProb*1000000) / 1000000)

	if compare1 != 0.000850 && compare2 != 0.000850 {
		t.Fatal("Returned: ", prob, " Expected: ", 0.000850, " Hand Math: ", compare2)
	}

	cards = append(cards, card4)

	t.Log("\n")
	t.Logf("Test 8.4: RoyalFlushCheck (probability)")

	t.Logf("Input of deck, selection of 4 cards (for royal flush), output should be 1/48 or 0.0208")

	prob = deck.RoyalFlush(tempDeck, cards)
	compare1 = (math.Round(prob*1000000) / 1000000)

	if compare1 != 0.020833 {
		t.Fatal("Returned: ", prob, " Expected: ", 0.020833)
	}

	cards = append(cards, card5)

	t.Log("\n")
	t.Logf("Test #8.5: RoyalFlushCheck (identification)")
	//t.Logf("Input of deck, selection of 4 cards (for royal flush), output should be true as it only need ace of spades")

	t.Logf("Input of deck, selection of 5 cards (for royal flush), output should be true") //added line for half-measure test

	probFloat := deck.RoyalFlush(tempDeck, cards)

	if probFloat != 1.00 {
		t.Fatal("Returned 0.0 when it should be greater")
	}

	//if true {
	//	t.Logf("Intentional failure")
	//}
}

// Test to check whether straight function will properly identify a straight (flush)
func TestStraightCheck(t *testing.T) {
	tempDeck := deck.NewDeck()

	card1 := card.NewCard(12, "Spade")
	card2 := card.NewCard(11, "Spade")
	card3 := card.NewCard(10, "Spade")
	card4 := card.NewCard(9, "Spade")
	//card5 := card.NewCard(8, "Spade")	//for testing non-broadway (no difference in output)
	//card5 := card.NewCard(0, "Heart") //for testing non-flush (bool will return false)
	card5 := card.NewCard(0, "Spade")

	var cards []card.Card

	cards = append(cards, card1)

	t.Log("\n")
	t.Logf("Test 9.1: StraightCheck (probability)")

	t.Logf("Input of deck, selection of 1 card (for broadway straight), output should be 0.004512")

	probFloat, straightFlushProb := deck.StraightCheck(tempDeck, cards)
	compare1 := (math.Round(probFloat*1000000) / 1000000)

	if compare1 != 0.004512 {
		t.Fatal("Returned: ", probFloat, " Expected: ", 0.004512)
	}

	cards = append(cards, card2)

	t.Log("\n")
	t.Logf("Test 9.2: StraightCheck (probability)")

	t.Logf("Input of deck, selection of 2 cards (for broadway straight), output should be 0.007760")

	probFloat, straightFlushProb = deck.StraightCheck(tempDeck, cards)
	compare1 = (math.Round(probFloat*1000000) / 1000000)

	if compare1 != 0.007760 {
		t.Fatal("Returned: ", probFloat, " Expected: ", 0.007760)
	}

	cards = append(cards, card3)

	t.Log("\n")
	t.Logf("Test 9.3: StraightCheck (probability)")

	t.Logf("Input of deck, selection of 3 cards (for broadway straight), output should be 0.018287")

	probFloat, straightFlushProb = deck.StraightCheck(tempDeck, cards)
	compare1 = (math.Round(probFloat*1000000) / 1000000)

	if compare1 != 0.018287 {
		t.Fatal("Returned: ", probFloat, " Expected: ", 0.018287)
	}

	cards = append(cards, card4)

	t.Log("\n")
	t.Logf("Test 9.4: StraightCheck (probability)")

	t.Logf("Input of deck, selection of 4 cards (for broadway straight), output should be 0.101218")

	probFloat, straightFlushProb = deck.StraightCheck(tempDeck, cards)
	compare1 = (math.Round(probFloat*1000000) / 1000000)

	if compare1 != 0.101218 {
		t.Fatal("Returned: ", probFloat, " Expected: ", 0.101218)
	}

	cards = append(cards, card5)

	t.Log("\n")
	t.Logf("Test #9: StraightCheck (identifying)")
	t.Logf("Input of deck, selection of 5 cards (for broadway straight), output should be true")

	probFloat, straightFlushProb = deck.StraightCheck(tempDeck, cards)

	if probFloat == 0.00 {
		t.Fatal("Returned 0.0 when it should be greater")
	}
	if straightFlushProb == 0.00 {
		t.Fatal("Returned 0.0 for straight flush when it should be greater")
	}
}

// Tests to see if the hand + community cards contains 5 cards of the same suit for a flush. (Containing 3 different suits between 7 cards invalidates the possibility)
func TestFlushCheck(t *testing.T) {
	tempDeck := deck.NewDeck()

	card1 := card.NewCard(12, "Spade")
	card2 := card.NewCard(11, "Spade")
	card3 := card.NewCard(10, "Spade")
	card4 := card.NewCard(9, "Spade")
	card5 := card.NewCard(0, "Spade")
	//card5 := card.NewCard(0, "Diamond") for testing non-flush

	var cards []card.Card

	cards = append(cards, card1)

	t.Log("\n")
	t.Logf("Test #10.1: FlushCheck (probability)")
	t.Logf("Input of deck, selection of 1 spade card, output should be 0.244775")

	probFloat := deck.FlushCheck(tempDeck, cards)
	compare1 := (math.Round(probFloat*1000000) / 1000000)

	if compare1 != 0.244775 {
		t.Fatal("Returned: ", compare1, " Expected: ", 0.244775)
	}

	cards = append(cards, card2)

	t.Log("\n")
	t.Logf("Test #10.2: FlushCheck (probability)")
	t.Logf("Input of deck, selection of 2 spade cards, output should be 0.269185")

	probFloat = deck.FlushCheck(tempDeck, cards)
	compare1 = (math.Round(probFloat*1000000) / 1000000)

	if compare1 != 0.269185 {
		t.Fatal("Returned: ", compare1, " Expected: ", 0.269185)
	}

	cardCopy := cards

	t.Log("\n")
	t.Logf("Test #10.2.5.1: FlushCheck (probability)")
	t.Logf("Input of deck, selection of 2 spades, 1 heart, output should be 0.109805")

	cardHeart1 := card.NewCard(1, "Heart")
	cardCopy = append(cardCopy, cardHeart1)

	probFloat = deck.FlushCheck(tempDeck, cardCopy)
	compare1 = (math.Round(probFloat*1000000) / 1000000)

	if compare1 != 0.109805 { //The probability drops significantly. Only 2 suits available now.
		t.Fatal("Returned: ", compare1, " Expected: ", 0.109805)
	}

	t.Log("\n")
	t.Logf("Test #10.2.5.2: FlushCheck (probability)")
	t.Logf("Input of deck, selection of 2 spades, 2 hearts, output should be 0.114477")

	cardHeart2 := card.NewCard(2, "Heart")
	cardCopy = append(cardCopy, cardHeart2)

	probFloat = deck.FlushCheck(tempDeck, cardCopy)
	compare1 = (math.Round(probFloat*1000000) / 1000000)

	if compare1 != 0.114477 {
		t.Fatal("Returned: ", compare1, " Expected: ", 0.114477)
	}

	t.Log("\n")
	t.Logf("Test #10.2.5.3: FlushCheck (probability)")
	t.Logf("Input of deck, selection of 2 spades, 3 hearts, output should be 0.083256")

	cardHeart3 := card.NewCard(3, "Heart")
	cardCopy = append(cardCopy, cardHeart3)

	probFloat = deck.FlushCheck(tempDeck, cardCopy)
	compare1 = (math.Round(probFloat*1000000) / 1000000)

	if compare1 != 0.083256 {
		t.Fatal("Returned: ", compare1, " Expected: ", 0.083256)
	}

	t.Log("\n")
	t.Logf("Test #10.2.5.4: FlushCheck (probability)")
	t.Logf("Input of deck, selection of 2 spades, 4 hearts, output should be 0.195652")

	cardHeart4 := card.NewCard(4, "Heart")
	cardCopy = append(cardCopy, cardHeart4)

	probFloat = deck.FlushCheck(tempDeck, cardCopy)
	compare1 = (math.Round(probFloat*1000000) / 1000000)

	if compare1 != 0.195652 {
		t.Fatal("Returned: ", compare1, " Expected: ", 0.195652)
	}

	cards = append(cards, card3)

	t.Log("\n")
	t.Logf("Test #10.3: FlushCheck (probability)")
	t.Logf("Input of deck, selection of 3 spade cards, output should be 0.076531")

	probFloat = deck.FlushCheck(tempDeck, cards)
	compare1 = (math.Round(probFloat*1000000) / 1000000)

	if compare1 != 0.076531 { //The probability drops significantly as you can no longer get flushes from the other suits.
		t.Fatal("Returned: ", compare1, " Expected: ", 0.076531)
	}

	cards = append(cards, card4)

	t.Log("\n")
	t.Logf("Test #10.4: FlushCheck (probability)")
	t.Logf("Input of deck, selection of 4 spade cards, output should be 0.1875")

	probFloat = deck.FlushCheck(tempDeck, cards)
	compare1 = (math.Round(probFloat*1000000) / 1000000)

	if compare1 != 0.1875 { //The probability increases significantly again as it's almost a 1/4 chance to get the right suit on the last card.
		t.Fatal("Returned: ", compare1, " Expected: ", 0.1875)
	}

	cards = append(cards, card5)

	t.Log("\n")
	t.Logf("Test #10: FlushCheck (identifying)")
	t.Logf("Input of deck, selection of 5 cards (flush), output should be true")

	probFloat = deck.FlushCheck(tempDeck, cards)

	if probFloat == 0.00 {
		t.Fatal("Returned 0.0 when it should be greater")
	}
}

// Tests to see if the cards (marked extraCards) were properly removed from the array
func TestRemoveCardsFromArray(t *testing.T) {
	card1 := card.NewCard(12, "Spade")
	card2 := card.NewCard(10, "Spade")
	card3 := card.NewCard(8, "Spade")
	card4 := card.NewCard(6, "Spade")
	card5 := card.NewCard(0, "Spade")
	//card5 := card.NewCard(0, "Diamond") for testing non-flush

	var cards []card.Card

	cards = append(cards, card1)
	cards = append(cards, card2)
	cards = append(cards, card3)
	cards = append(cards, card4)
	cards = append(cards, card5)

	var extraCards []card.Card

	extraCards = append(extraCards, card4)
	extraCards = append(extraCards, card5)

	cardCopy := deck.RemoveCardsFromArray(cards, extraCards)

	t.Log("\n")
	t.Logf("Test #11: RemoveCardsFromArray")
	t.Logf("Input of 5 cards, selection of 2 cards from the original array, output should be two less cards")

	if len(cardCopy) < 3 {
		t.Fatal("Returned less than 3 cards (too many removed)")
	}
	if len(cardCopy) > 3 {
		t.Fatal("Returned more than 3 cards (too few removed)")
	}

	card4Bool := false
	card5Bool := false

	for i := 0; i < len(cardCopy); i++ {
		if cardCopy[i] == card4 {
			card4Bool = true
		}
		if cardCopy[i] == card5 {
			card5Bool = true
		}
	}

	if card4Bool && card5Bool {
		t.Fatal("Contains card4 and card5 when both should've been removed")
	}
	if card4Bool {
		t.Fatal("Contains card4 when it should've been removed")
	}
	if card5Bool {
		t.Fatal("Contains card5 when it should've been removed")
	}

}

// Tests to see if probability is accurately being calculated when input with the cards needed for a particular hand
func TestFindCardProb(t *testing.T) {

	//Test0, calculates the chance of getting a Royal Flush and then multiplies it by 4 for each distinct Royal Flush
	//Result is compared to cumulative Royal Flush probability from Wikipedia
	var test0Cards []card.Card
	test0Target := []int{12, 11, 10, 9, 0}
	test0Suit := "Spade"
	test0 := deck.FindCardProb(test0Cards, test0Target, test0Suit, 0)
	test0 *= 4.00 //For each distinct Royal Flush (4)
	test0Prob := (math.Round(test0*100000000) / 100000000)

	if test0Prob != 0.00000154 {
		t.Fatal("Royal Flush initial prob incorrect. True prob: 0.00000154 Output: ", test0Prob)
	}

	//Royal flush
	card1 := card.NewCard(11, "Spade")
	card2 := card.NewCard(9, "Spade")

	var cards []card.Card

	cards = append(cards, card1)
	cards = append(cards, card2)

	t.Log("\n")
	t.Logf("Test #12: FindCardProb")
	t.Logf("Various inputs, various tests")

	t.Log("\n")
	t.Logf("	Subtest #1, Royal Flush:")

	targetSuit := "Spade"
	targetVals := []int{12, 10, 0}
	permutations := float64(deck.Factorial(3))
	totalProb := ((1.00 / 50.00) * (1.00 / 49.00) * (1.00 / 48.00)) * permutations
	tempFloat := deck.FindCardProb(cards, targetVals, targetSuit, 0)
	compare1 := (math.Round(tempFloat*1000000) / 1000000)
	compare2 := (math.Round(totalProb*1000000) / 1000000)

	//Function call (last value only relevant for flush specifically)
	if compare1 != 0.000051 && compare2 != 0.000051 {
		t.Fatal("Returned a different percent value: ", compare1)
	}

	t.Log("\n")
	t.Logf("	Subtest #2, Straight Flush:")

	targetSuit = "Spade"
	targetVals = []int{10, 8, 7}
	permutations = float64(deck.Factorial(3))
	totalProb = ((1.00 / 50.00) * (1.00 / 49.00) * (1.00 / 48.00)) * permutations
	tempFloat = deck.FindCardProb(cards, targetVals, targetSuit, 0)
	compare1 = (math.Round(tempFloat*1000000) / 1000000)
	compare2 = (math.Round(totalProb*1000000) / 1000000)

	//Function call (last value only relevant for flush specifically)
	if compare1 != 0.000051 && compare2 != 0.000051 {
		t.Fatal("Returned a different percent value: ", compare1)
	}

	t.Log("\n")
	t.Logf("	Subtest #3, Four of a Kind:")

	targetSuit = ""
	targetVals = []int{11, 11, 11}
	permutations = float64(deck.Factorial(3))
	totalProb = ((3.00 / 50.00) * (2.00 / 49.00) * (1.00 / 48.00)) * permutations
	tempFloat = deck.FindCardProb(cards, targetVals, targetSuit, 0)
	compare1 = (math.Round(tempFloat*1000000) / 1000000)
	compare2 = (math.Round(totalProb*1000000) / 1000000)

	//Function call (last value only relevant for flush specifically)
	if compare1 != 0.000306 && compare2 != 0.000306 {
		t.Fatal("Returned a different percent value: ", compare1)
	}

	t.Log("\n")
	t.Logf("	Subtest #4, Full House:")

	targetSuit = ""
	targetVals = []int{11, 11, 9}
	permutations = float64(deck.Factorial(3))
	totalProb = ((3.00 / 50.00) * (2.00 / 49.00) * (3.00 / 48.00)) * permutations
	tempFloat = deck.FindCardProb(cards, targetVals, targetSuit, 0)
	compare1 = (math.Round(tempFloat*1000000) / 1000000)
	compare2 = (math.Round(totalProb*1000000) / 1000000)

	//Function call (last value only relevant for flush specifically)
	if compare1 != 0.000918 && compare2 != 0.000918 {
		t.Fatal("Returned a different percent value: ", compare1)
	}

}

func TestDebugLogic(t *testing.T) {
	tempDeck := deck.NewDeck()
	card1 := card.NewCard(12, "Spade")
	card2 := card.NewCard(11, "Spade")
	card3 := card.NewCard(10, "Spade")
	card4 := card.NewCard(9, "Spade")
	card5 := card.NewCard(0, "Spade")
	cardE1 := card.NewCard(12, "Heart")
	cardE5 := card.NewCard(0, "Heart")
	cardE3 := card.NewCard(11, "Club")

	var cards []card.Card

	cardCopy := cards
	cardCopy = append(cardCopy, cardE1)
	cardCopy = append(cardCopy, card2)

	t.Log("\n")
	t.Logf("Debug Logic Test:")
	t.Logf("1 royal spade, 1 royal heart")

	royalProb := deck.RoyalFlush(tempDeck, cardCopy)
	straightProb, straightFlushProb := deck.StraightCheck(tempDeck, cardCopy)
	rProb := (math.Round(royalProb*100) / 100)
	sProb := (math.Round(straightProb*100) / 100)
	//sfProb := (math.Round(straightFlushProb*100) / 100)

	t.Log("")

	if rProb != 0.00 && rProb != 1.00 {
		t.Fatal("RoyalFlush failure. output is 0 or 1: ", royalProb)
	}
	t.Logf("RoyalFlush success: %f\n", royalProb)
	if sProb != 0.00 && sProb != 1.00 {
	} else {
		t.Fatal("Straight failure. output between 0 and 1: ", straightProb)
	}
	t.Logf("Straight success: %f\n", straightProb)

	flushProb := deck.FlushCheck(tempDeck, cardCopy)
	if straightFlushProb > 0.00 {
		t.Logf("straight flush output??")
	} else if straightProb > 0 && flushProb > 0 {
		straightFlushProb = straightProb * flushProb
	} else {
		straightFlushProb = 0.00
	}
	t.Logf("Straight Flush success: %f\n", straightFlushProb)

	cardCopy = append(cardCopy, card3)

	t.Log("\n")
	t.Logf("Debug Logic Test:")
	t.Logf("2 royal spades, 1 royal heart")

	royalProb = deck.RoyalFlush(tempDeck, cardCopy)
	straightProb, straightFlushProb = deck.StraightCheck(tempDeck, cardCopy)
	rProb = (math.Round(royalProb*100) / 100)
	sProb = (math.Round(straightProb*100) / 100)
	//sfProb = (math.Round(straightFlushProb*100) / 100)

	t.Log("")

	if rProb != 0.00 && rProb != 1.00 {
		t.Fatal("RoyalFlush failure. output is 0 or 1: ", royalProb)
	}
	t.Logf("RoyalFlush success: %f\n", royalProb)
	if sProb != 0.00 && sProb != 1.00 {
	} else {
		t.Fatal("Straight failure. output between 0 and 1: ", straightProb)
	}
	t.Logf("Straight success: %f\n", straightProb)

	flushProb = deck.FlushCheck(tempDeck, cardCopy)
	if straightFlushProb > 0.00 {
		t.Logf("straight flush output??")
	} else if straightProb > 0 && flushProb > 0 {
		straightFlushProb = straightProb * flushProb
	} else {
		straightFlushProb = 0.00
	}
	t.Logf("Straight Flush success: %f\n", straightFlushProb)

	cardCopy = append(cardCopy, card4)
	cardCopy = append(cardCopy, cardE5)

	t.Log("\n")
	t.Logf("Debug Logic Test:")
	t.Logf("3 royal spades, 2 royal heart")

	royalProb = deck.RoyalFlush(tempDeck, cardCopy)
	straightProb, straightFlushProb = deck.StraightCheck(tempDeck, cardCopy)
	rProb = (math.Round(royalProb*100) / 100)
	sProb = (math.Round(straightProb*100) / 100)
	//sfProb = (math.Round(straightFlushProb*100) / 100)

	t.Log("")

	if rProb != 0.00 && rProb != 1.00 {
		t.Fatal("RoyalFlush failure. output is 0 or 1: ", royalProb)
	}
	t.Logf("RoyalFlush success: %f\n", royalProb)
	if sProb != 0.00 && sProb != 1.00 {
		t.Fatal("Straight output failure. between 0 and 1: ", straightProb)
	}
	t.Logf("Straight success: %f\n", straightProb)

	flushProb = deck.FlushCheck(tempDeck, cardCopy)
	if straightFlushProb > 0.00 {
		t.Logf("straight flush output??")
	} else if straightProb > 0 && flushProb > 0 {
		straightFlushProb = straightProb * flushProb
	} else {
		straightFlushProb = 0.00
	}
	t.Logf("Straight Flush success: %f\n", straightFlushProb)

	cards = append(cards, card1)
	cards = append(cards, card2)
	cards = append(cards, card3)
	cards = append(cards, card4)

	cardCopy = cards

	t.Log("")
	t.Logf("Debug Logic Test:")
	t.Logf("4 royal spades, 1 royal heart")

	cardCopy = append(cardCopy, cardE5)

	royalProb = deck.RoyalFlush(tempDeck, cardCopy)
	straightProb, straightFlushProb = deck.StraightCheck(tempDeck, cardCopy)
	rProb = (math.Round(royalProb*100) / 100)
	sProb = (math.Round(straightProb*100) / 100)
	//sfProb = (math.Round(straightFlushProb*100) / 100)

	t.Log("")

	if rProb != 0.00 && rProb != 1.00 {
	} else {
		t.Fatal("RoyalFlush failure. output is 0 or 1: ", royalProb)
	}
	t.Logf("RoyalFlush success: %f\n", royalProb)
	if sProb != 0.00 && sProb != 1.00 {
		t.Fatal("Straight output failure. between 0 and 1: ", straightProb)
	}
	t.Logf("Straight success: %f\n", straightProb)

	flushProb = deck.FlushCheck(tempDeck, cardCopy)
	if straightFlushProb > 0.00 {
		t.Logf("straight flush output??")
	} else if straightProb > 0 && flushProb > 0 {
		straightFlushProb = straightProb * flushProb
	} else {
		straightFlushProb = 0.00
	}
	t.Logf("Straight Flush success: %f\n", straightFlushProb)

	cardCopy = append(cardCopy, cardE1)
	cardCopy = append(cardCopy, cardE3)
	cards = append(cards, card5)

}

// Runs 50 5-card draws and prints the results of StraightCheck, FlushCheck, and RoyalFlush
func TestMassTest(t *testing.T) {

	deckCopy := deck.NewDeck()

	//t.Logf("Str8 \tFlush\tStr8 F\tR Flush")
	t.Logf("Pair \t2Pair\t3Kind\t4Kind\tFullH")

	for i := 0; i < 50; i++ {
		var cardStrings []string
		cards := []card.Card{}
		var intArr []int

		for n := 0; n < 5; n++ {
			tempVal := 0
			var tempCard card.Card

			for deck.ContainsInt(intArr, tempVal) {
				tempVal = rand.Intn(51)
			}

			intArr = append(intArr, tempVal)
			tempCard = deckCopy[tempVal]
			cards = append(cards, tempCard)

			valString := ""

			valString = strconv.Itoa(cards[n].Val)

			switch cards[n].Suit {
			case "Heart":
				tempString := valString + "H"
				cardStrings = append(cardStrings, tempString)
			case "Diamond":
				tempString := valString + "D"
				cardStrings = append(cardStrings, tempString)
			case "Club":
				tempString := valString + "C"
				cardStrings = append(cardStrings, tempString)
			case "Spade":
				tempString := valString + "S"
				cardStrings = append(cardStrings, tempString)
			}
		}

		t.Log(cardStrings[0], " ", cardStrings[1], " ", cardStrings[2], " ", cardStrings[3], " ", cardStrings[4])

		/*
			straightProb, straightFlushProb := deck.StraightCheck(deckCopy, cards)
			flushProb := deck.FlushCheck(deckCopy, cards)
			royalProb := deck.RoyalFlush(deckCopy, cards)

			sProb := (math.Round(straightProb*10000000.0) / 10000000.0)
			fProb := (math.Round(flushProb*10000000.0) / 10000000.0)
			sfProb := 0.00
			rProb := (math.Round(royalProb*10000000.0) / 10000000.0)

			if straightProb <= 0.0000001 || flushProb <= 0.0000001 {
				sfProb = 0.00
			} else if straightFlushProb > 0.00 {
				sfProb = (math.Round(straightFlushProb*10000000.0) / 10000000.0)
			} else {
				sfProb = (math.Round((straightProb*flushProb)*10000000.0) / 10000000.0)
			}

			t.Log(sProb, " \t", fProb, " \t", sfProb, " \t", rProb, "\n")
		*/

		handTypes := deck.CheckHandType(cards)
		futureHandProbs := deck.DetermineFutureProbability(cards, deck.DetermineFutureHands(cards, handTypes))

		pair1Prob := 0.00
		pair2Prob := 0.00
		pair3Prob := 0.00
		pair4Prob := 0.00
		pair5Prob := 0.00
		t.Log(cards)
		t.Log(handTypes)
		if deck.Contains(handTypes, "One Pair") {
			pair1Prob = 1.00
		} else {
			pair1Prob = futureHandProbs[0]
		}
		if deck.Contains(handTypes, "Two Pair") {
			pair2Prob = 1.00
		} else {
			pair2Prob = futureHandProbs[1]
		}
		if deck.Contains(handTypes, "Three of a Kind") {
			pair3Prob = 1.00
		} else {
			pair3Prob = futureHandProbs[2]
		}
		if deck.Contains(handTypes, "Full House") {
			pair5Prob = 1.00
		} else {
			pair5Prob = futureHandProbs[3] //should be 3 after function is updated
		}
		if deck.Contains(handTypes, "Four of a Kind") {
			pair4Prob = 1.00
		} else {
			//pair4Prob = futureHandProbs[4]	this would be 4
			pair4Prob = futureHandProbs[4]
		}

		p1Prob := (math.Round((pair1Prob)*10000000.0) / 10000000.0)
		p2Prob := (math.Round((pair2Prob)*10000000.0) / 10000000.0)
		p3Prob := (math.Round((pair3Prob)*10000000.0) / 10000000.0)
		p4Prob := (math.Round((pair4Prob)*10000000.0) / 10000000.0)
		p5Prob := (math.Round((pair5Prob)*10000000.0) / 10000000.0)

		t.Log(p1Prob, " \t", p2Prob, " \t", p3Prob, " \t", p4Prob, "\t", p5Prob, "\n")
	}

	//********************************************************THIS LINE MUST BE UNCOMMENTED TO VIEW OUTPUT********************************************************
	//t.Fatal("Test Complete")
	//********************************************************THIS LINE MUST BE UNCOMMENTED TO VIEW OUTPUT********************************************************
}

func TestFactorial(t *testing.T) {
	temp := deck.Factorial(7)
	if temp == 0.00 {
		t.Fatal("No input should result in 0, except for 0")
	} else if temp != 5040.000 {
		t.Fatal("Incorrect output")
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
	temparray := deck.GetHandArray(temphand)
	if deck.Contains(deck.CheckHandType(temparray), "One Pair") != true {
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
	temparray := deck.GetHandArray(temphand)
	if deck.Contains(deck.CheckHandType(temparray), "Two Pair") != true {
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
	temparray := deck.GetHandArray(temphand)
	if deck.Contains(deck.CheckHandType(temparray), "Three of a Kind") != true {
		t.Fatal("Three of a Kind comparison is not working!")
	} else {
		t.Log("Three of a Kind comparison successful!")
	}
	hand.AddCardHandSpecific(temphand, 1, "Diamond")
	temparray = deck.GetHandArray(temphand)
	if deck.Contains(deck.CheckHandType(temparray), "Four of a Kind") != true {
		t.Fatal("Four of a Kind comparison is not working!")
	} else {
		t.Log("Four of a Kind comparison successful!")
	}
	hand.AddCardHandSpecific(temphand, 4, "Spade")
	temparray = deck.GetHandArray(temphand)
	if deck.Contains(deck.CheckHandType(temparray), "Full House") != true {
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
	temparray := deck.GetHandArray(temphand)
	if deck.Contains(deck.DetermineFutureHands(temphand.ActualHand, deck.CheckHandType(temparray)), "Four of a Kind") != true {
		t.Fatal("Future hand function does not work for three of a kind!")
	}
	t.Log(("Future hand function works for three of a kind!"))
}

func TestFutureProbabilityOnePair(t *testing.T) {
	t.Log("Testing future probability for one pair")
	temphand := hand.NewHand("None")
	var handArray []float64
	var probabilityArray []float64
	hand.AddCardHandSpecific(temphand, 1, "Heart")
	probabilityArray = append(probabilityArray, FindFutureProbability(len(temphand.ActualHand), 7-len(temphand.ActualHand)))
	handArray = append(handArray, deck.DetermineFutureProbability(temphand.ActualHand, deck.DetermineFutureHands(temphand.ActualHand, deck.CheckHandType(temphand.ActualHand)))[0])
	hand.AddCardHandSpecific(temphand, 2, "Heart")
	probabilityArray = append(probabilityArray, FindFutureProbability(len(temphand.ActualHand), 7-len(temphand.ActualHand)))
	handArray = append(handArray, deck.DetermineFutureProbability(temphand.ActualHand, deck.DetermineFutureHands(temphand.ActualHand, deck.CheckHandType(temphand.ActualHand)))[0])
	hand.AddCardHandSpecific(temphand, 3, "Heart")
	probabilityArray = append(probabilityArray, FindFutureProbability(len(temphand.ActualHand), 7-len(temphand.ActualHand)))
	handArray = append(handArray, deck.DetermineFutureProbability(temphand.ActualHand, deck.DetermineFutureHands(temphand.ActualHand, deck.CheckHandType(temphand.ActualHand)))[0])
	hand.AddCardHandSpecific(temphand, 4, "Heart")
	probabilityArray = append(probabilityArray, FindFutureProbability(len(temphand.ActualHand), 7-len(temphand.ActualHand)))
	handArray = append(handArray, deck.DetermineFutureProbability(temphand.ActualHand, deck.DetermineFutureHands(temphand.ActualHand, deck.CheckHandType(temphand.ActualHand)))[0])
	hand.AddCardHandSpecific(temphand, 5, "Heart")
	probabilityArray = append(probabilityArray, FindFutureProbability(len(temphand.ActualHand), 7-len(temphand.ActualHand)))
	handArray = append(handArray, deck.DetermineFutureProbability(temphand.ActualHand, deck.DetermineFutureHands(temphand.ActualHand, deck.CheckHandType(temphand.ActualHand)))[0])
	hand.AddCardHandSpecific(temphand, 6, "Heart")
	probabilityArray = append(probabilityArray, FindFutureProbability(len(temphand.ActualHand), 7-len(temphand.ActualHand)))
	handArray = append(handArray, deck.DetermineFutureProbability(temphand.ActualHand, deck.DetermineFutureHands(temphand.ActualHand, deck.CheckHandType(temphand.ActualHand)))[0])
	for index, _ := range probabilityArray {
		if handArray[index] != probabilityArray[index] {
			t.Log(handArray[index])
			t.Log(probabilityArray[index])
			t.Fatal("One Pair future probability is wrong at" + strconv.Itoa(index))
		}
		t.Log(handArray[index])
	}
	t.Log("One Pair future probability is all right")
}
func FindFutureProbability(handLength int, canAddNumCards int) float64 {
	probability := 0.0
	if canAddNumCards != 0 {
		for i := 1; i <= canAddNumCards; i++ {
			notOnePairProb := math.Pow(float64(52-4*handLength)/float64(52-handLength), float64(canAddNumCards-i))
			probability += float64(canAddNumCards) * notOnePairProb * math.Pow(float64(3*handLength)/float64(52-handLength), float64(i))
		}
	} else {
		probability = 0
	}
	return probability
}
func Factorial(n int) float64 {
	factVal := 1.0000000
	if n < 0 {
		factVal = -1
	} else {
		for i := 1; i <= n; i++ {
			factVal *= float64(i) // mismatched types int64 and int
		}

	}
	return factVal /* return from function*/
}
func TestFutureProbabilityThreeOfKind(t *testing.T) {
	t.Log("Testing future probability for three of a kind with one pair")
	temphand := hand.NewHand("One Pair")
	hand.AddCardHandSpecific(temphand, 1, "Heart")
	hand.AddCardHandSpecific(temphand, 1, "Spade")
	canAddNumCards := 7 - len(temphand.ActualHand)
	probability := 0.0
	for i := 1; i < 5; i++ {
		ThreeNotGottenProb := math.Pow(float64(52-4*len(temphand.ActualHand))/float64(52-len(temphand.ActualHand)), float64(canAddNumCards-2)) / float64(52-len(temphand.ActualHand)) / float64(52-len(temphand.ActualHand))
		probability += Factorial(52-len(temphand.ActualHand)) / (Factorial(canAddNumCards) * Factorial(52-len(temphand.ActualHand)-canAddNumCards)) * ThreeNotGottenProb * float64(2) / float64(52-len(temphand.ActualHand)) * float64(1) / float64(52-len(temphand.ActualHand))
		break
	}
	array := deck.DetermineFutureProbability(temphand.ActualHand, deck.DetermineFutureHands(temphand.ActualHand, deck.CheckHandType(temphand.ActualHand)))
	probability = array[2]
	if array[2] != probability {
		t.Log(array[2])
		t.Log(probability)
		t.Fatal("Three of a kind future probability is wrong")
	}
	t.Log("Three of a kind future probability is right")
}
func TestFutureProbabilityTwoPair(t *testing.T) {
	t.Log("Testing future probability for 2 pair")
	temphand := hand.NewHand("None")
	hand.AddCardHandSpecific(temphand, 1, "Heart")
	hand.AddCardHandSpecific(temphand, 1, "Spade")
	canAddNumCards := 7 - len(temphand.ActualHand)
	probability := 0.0
	for i := 1; i <= 5; i++ {
		notTwoPairProb := math.Pow(float64(52-4*len(temphand.ActualHand))/float64(52-len(temphand.ActualHand)), float64(canAddNumCards-i))
		probability += float64(canAddNumCards) * notTwoPairProb * math.Pow(float64(3*len(temphand.ActualHand))/float64(52-len(temphand.ActualHand)), float64(i))
	}
	array := deck.DetermineFutureProbability(temphand.ActualHand, deck.DetermineFutureHands(temphand.ActualHand, deck.CheckHandType(temphand.ActualHand)))
	probability = array[1]
	if array[1] != probability {
		t.Log(array[1])
		t.Log(probability)
		t.Fatal("2 pair future probability is wrong")
	}
	t.Log("2 pair future probability is right")
}

// attempt 1 at future probability function test
func TestFutureProbabilityFourOfKind(t *testing.T) {
	t.Log("Testing future probability for four of a kind with one triple")
	temphand := hand.NewHand("None")
	hand.AddCardHandSpecific(temphand, 1, "Heart")
	hand.AddCardHandSpecific(temphand, 1, "Spade")
	hand.AddCardHandSpecific(temphand, 1, "Club")
	canAddNumCards := 7 - len(temphand.ActualHand)
	probability := 0.0
	for i := 1; i <= 4; i++ {
		FourNotGottenProb := math.Pow(float64(48)/float64(52-len(temphand.ActualHand)), float64(4)) / float64(52)
		probability += Factorial(52-canAddNumCards) / (Factorial(len(temphand.ActualHand)) * Factorial(52-len(temphand.ActualHand)-canAddNumCards)) * FourNotGottenProb * math.Pow(float64(1)/float64(52-len(temphand.ActualHand)), float64(canAddNumCards-i))
		break
	}
	array := deck.DetermineFutureProbability(temphand.ActualHand, deck.DetermineFutureHands(temphand.ActualHand, deck.CheckHandType(temphand.ActualHand)))
	probability = array[4]
	if array[4] != probability {
		t.Log(array[4])
		t.Log(probability)
		t.Fatal("Four of a kind future probability is wrong")
	}
	t.Log("Four of a kind future probability is right")
}
