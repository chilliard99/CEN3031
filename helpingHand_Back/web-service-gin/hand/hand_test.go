package hand_test

import (
	"example/web-service-gin/deck"
	"example/web-service-gin/hand"
	"strconv"
	"strings"
	"testing"
)

// Attempts to see if hand type changes after inserting one pair of cards
func TestOnePairCheck(t *testing.T) {
	t.Log("Testing one pair functionality")
	temphand := hand.NewHand("None")
	hand.AddCardHandSpecific(temphand, 1, "Heart")
	hand.AddCardHandSpecific(temphand, 2, "Heart")
	hand.AddCardHandSpecific(temphand, 1, "Club")
	hand.AddCardHandSpecific(temphand, 3, "Spade")
	hand.AddCardHandSpecific(temphand, 4, "Diamond")
	if strings.Compare(deck.CheckHandType(temphand), "One Pair") != 0 {
		t.Fatal("One Pair comparison is not working!")
	} else {
		t.Log("One Pair comparison successful!")
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

// Attempts to add more than 5 cards to the hand to check if it returns the error message.
func TestHandAddOverflow(t *testing.T) {
	t.Log("Testing adding more than 5 cards to hand")
	temphand := hand.NewHand("None")
	//First 5 adds should work without errors
	hand.AddCardHandRandom(temphand)
	hand.AddCardHandRandom(temphand)
	hand.AddCardHandRandom(temphand)
	hand.AddCardHandRandom(temphand)
	hand.AddCardHandRandom(temphand)
	//Seeing if the returned string after attempting to add a 6th card is correct
	if strings.Compare(hand.AddCardHandRandom(temphand), "adding random card: length of hand is already 5") != 0 {
		t.Fatal("Card should not have been added")
	}
}
