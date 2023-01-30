package deck_test

import (
	"example/web-service-gin/deck"
	"testing"
)

func TestDeck(t *testing.T) {
	if deck.ReturnDeckSize() != 52 {
		t.Fatal("Deck not full!")
	}
}
