package hand_test

import (
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
	//First 7 adds should work without errors
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
