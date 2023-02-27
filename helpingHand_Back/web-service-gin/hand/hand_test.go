package hand_test

import (
	"example/web-service-gin/deck"
	"example/web-service-gin/hand"
	"strconv"
	"strings"
	"testing"
)

// check if string slice already contains item
// from https://freshman.tech/snippets/go/check-if-slice-contains-element/
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
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
	if contains(deck.CheckHandType(temphand), "One Pair") != true {
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
	if contains(deck.CheckHandType(temphand), "Two Pair") != true {
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
	if contains(deck.CheckHandType(temphand), "Three of a Kind") != true {
		t.Fatal("Three of a Kind comparison is not working!")
	} else {
		t.Log("Three of a Kind comparison successful!")
	}
	hand.AddCardHandSpecific(temphand, 1, "Diamond")
	if contains(deck.CheckHandType(temphand), "Four of a Kind") != true {
		t.Fatal("Four of a Kind comparison is not working!")
	} else {
		t.Log("Four of a Kind comparison successful!")
	}
	hand.AddCardHandSpecific(temphand, 4, "Spade")
	if contains(deck.CheckHandType(temphand), "Full House") != true {
		t.Fatal("Full House comparison is not working!")
	} else {
		t.Log("Full House comparison successful!")
	}
}

// Attempts to check 5 individual card values and suits after adding them to the hand
func TestHandCardValues(t *testing.T) {
	t.Log("Testing insertion of specific card values")
	temphand := hand.NewHand("None")
	//Adding and then checking all 5 indexes in the hand for correct suits/values
	if strings.Compare(hand.AddCardHandSpecific(temphand, 1, "Heart"), "adding specific card: successful") != 0 {
		t.Fatal("First card was not added")
	}
	if strings.Compare(hand.AddCardHandSpecific(temphand, 2, "Heart"), "adding specific card: successful") != 0 {
		t.Fatal("Second card was not added")
	}
	if strings.Compare(hand.AddCardHandSpecific(temphand, 12, "Diamond"), "adding specific card: successful") != 0 {
		t.Fatal("Third card was not added")
	}
	if strings.Compare(hand.AddCardHandSpecific(temphand, 10, "Spade"), "adding specific card: successful") != 0 {
		t.Fatal("Fourth card was not added")
	}
	if strings.Compare(hand.AddCardHandSpecific(temphand, 3, "Club"), "adding specific card: successful") != 0 {
		t.Fatal("Fifth card was not added")
	}
	for i := 0; i < 5; i++ {
		if temphand.ActualHand[i].Val != hand.CheckCardIndex(temphand, i).Val {
			t.Fatal("Index " + strconv.Itoa(i) + ": Card value is incorrect, should be " + strconv.Itoa(temphand.ActualHand[i].Val) + " but is " + strconv.Itoa(hand.CheckCardIndex(temphand, i).Val))
		}
		if strings.Compare(temphand.ActualHand[i].Suit, hand.CheckCardIndex(temphand, i).Suit) != 0 {
			t.Fatal("Index " + strconv.Itoa(i) + ": Card suit is incorrect, should be " + temphand.ActualHand[i].Suit + " but is " + hand.CheckCardIndex(temphand, i).Suit)
		}
	}
}

// Attempts to add more than 7 cards to the hand to check if it returns the error message.
func TestHandAddOverflow(t *testing.T) {
	t.Log("Testing adding more than 7 cards to hand")
	temphand := hand.NewHand("None")
	//First 5 adds should work without errors
	hand.AddCardHandRandom(temphand)
	hand.AddCardHandRandom(temphand)
	hand.AddCardHandRandom(temphand)
	hand.AddCardHandRandom(temphand)
	hand.AddCardHandRandom(temphand)
	hand.AddCardHandRandom(temphand)
	hand.AddCardHandRandom(temphand)
	//Seeing if the returned string after attempting to add a 8th card is correct
	if strings.Compare(hand.AddCardHandRandom(temphand), "adding random card: length of hand is already 7") != 0 {
		t.Fatal("Card should not have been added")
	}
}
