package card

// Card structure
type Card struct {
	Val   int    `json:"Val"`
	Suit  string `json:"Suit"`
	Index int    `json:"Index"`
}

func NewCard(value int, suit string) Card {
	var card Card
	card.Val = value
	card.Suit = suit
	card.Index = 0
	return card
}

/*
Card Indexes (only applicable to user hand)
	1 = hole 1
	2 = hole 2
	3 = community 1
	4 = community 2
	5 = community 3
	6 = community 4
	7 = community 5
*/

// Uses value and suit to return the name of the Card (i.e. "King of Clubs")
func GetCardName(card Card) string {
	value := ""

	//Simple switch case to translate 0-12 to proper card names
	switch card.Val {
	case 0:
		value = "Ace"
	case 1:
		value = "Two"
	case 2:
		value = "Three"
	case 3:
		value = "Four"
	case 4:
		value = "Five"
	case 5:
		value = "Six"
	case 6:
		value = "Seven"
	case 7:
		value = "Eight"
	case 8:
		value = "Nine"
	case 9:
		value = "Ten"
	case 10:
		value = "Jack"
	case 11:
		value = "Queen"
	case 12:
		value = "King"
	default:
		value = "error"
	}

	return value + " of " + card.Suit + "s"
}
