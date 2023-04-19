package deck

import (
	//imports card.go as c to prevent redundant card.Card or card.NewCard(int, string) every time a card is created
	c "example/web-service-gin/card"
	//imports hand.go as h to allow hand functions to be called
	h "example/web-service-gin/hand"
	"fmt"
	"math"
)

// Potential structure to return probability function to front end

type Getter interface {
	GetAll() []HandProb
}

type HandProb struct {
	Handname string  `json:"Handname"`
	Prob     float64 `json:"Prob"`
}

type Probabilitys struct {
	ProbList []HandProb
}

func New() *Probabilitys {
	return &Probabilitys{
		//Indexes:
		//0: High Card, 1: One Pair, 2: Two Pair, 3: Three of a Kind, 4: Straight, 5: Flush, 6: Full House, 7: Four of a Kind, 8: Straight Flush, 9: Royal Flush
		ProbList: []HandProb{{Handname: "High Card", Prob: 0.00}, {Handname: "One Pair", Prob: 0.00}, {Handname: "Two Pair", Prob: 0.00}, {Handname: "Three of a Kind", Prob: 0.00},
			{Handname: "Straight", Prob: 0.00}, {Handname: "Flush", Prob: 0.00}, {Handname: "Full House", Prob: 0.00}, {Handname: "Four of a Kind", Prob: 0.00},
			{Handname: "Straight Flush", Prob: 0.00}, {Handname: "Royal Flush", Prob: 0.00}},
	}
}

func (r *Probabilitys) GetAll() []HandProb {
	return r.ProbList
}

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

// Check if int slice already contains items
// Copied above variation with ints
func ContainsInt(n []int, num int) bool {
	for _, v := range n {
		if v == num {
			return true
		}
	}

	return false
}

// Define a deck as an array of cards
type Deck []c.Card

// Returns RoyalFlush output when given cards from the frontend
func UpdateProb(cards_ []c.Card, deck Deck, currUserProb []HandProb) {

	cards := []c.Card{}

	for i := 0; i < len(cards_); i++ {
		if cards_[i].Suit != "" {
			//fmt.Println("Card #", i, ": val=", cards_[i].Val, " suit=", cards_[i].Suit)
			cards = append(cards, cards_[i])
		}
	}

	//fmt.Println(len(cards))

	handTypes := CheckHandType(cards)
	futureHandProbs := DetermineFutureProbability(cards, DetermineFutureHands(cards, handTypes))
	deckCopy := RemoveCards(deck, cards)

	straightProb, straightFlushProb := StraightCheck(deckCopy, cards)
	flushProb := FlushCheck(deckCopy, cards)

	currUserProb[0].Prob = HighCard(deckCopy, cards)
	if Contains(handTypes, "One Pair") {
		currUserProb[1].Prob = 1.00
	} else {
		currUserProb[1].Prob = futureHandProbs[0]
	}
	if Contains(handTypes, "Two Pair") {
		currUserProb[2].Prob = 1.00
	} else {
		currUserProb[2].Prob = futureHandProbs[1]
	}
	if Contains(handTypes, "Three of a Kind") {
		currUserProb[3].Prob = 1.00
	} else {
		currUserProb[3].Prob = futureHandProbs[2]
	}
	currUserProb[4].Prob = straightProb
	currUserProb[5].Prob = flushProb

	if Contains(handTypes, "Full House") {
		currUserProb[6].Prob = 1.00
	} else {
		currUserProb[6].Prob = futureHandProbs[3] //should be 3 after function is updated
	}
	if Contains(handTypes, "Four of a Kind") {
		currUserProb[7].Prob = 1.00
	} else {
		currUserProb[7].Prob = futureHandProbs[4]
	}

	if flushProb <= 0.0000001 || straightProb <= 0.0000001 {
		currUserProb[8].Prob = 0.00
	} else if straightFlushProb > 0.00 {
		currUserProb[8].Prob = straightFlushProb
	} else {
		currUserProb[8].Prob = straightProb * flushProb
	}
	currUserProb[9].Prob = RoyalFlush(deckCopy, cards)

	//fmt.Println("Bool result: ", royalBoolean, " Probability: ", royalProb)
}

// Factorial function taken from:
// https://www.golangprograms.com/go-program-to-find-factorial-of-a-number.html
func Factorial(n int) float64 {
	factVal := 1.0000000
	if n < 0 {
		fmt.Print("Factorial of negative number doesn't exist.")
	} else {
		for i := 1; i <= n; i++ {
			factVal *= float64(i) // mismatched types int64 and int
		}

	}
	return factVal /* return from function*/
}

func FindCardProb(cards []c.Card, targetVals []int, targetSuit string, numSuitNeed int) float64 {
	deck := NewDeck()
	deckCopy := RemoveCards(deck, cards)
	cardCount := len(cards)
	numNeeded := len(targetVals)
	numToDraw := 7 - cardCount
	totalProb := 1.00
	var indProbs []float64

	//Default 0.00 if not enough cards will be drawn
	if numNeeded != 0 && numNeeded > numToDraw {
		return 0.00
	}

	//For flush only, where specific values are not needed.
	if numNeeded == 0 {
		for i := 0; i < numSuitNeed; i++ {
			validCardCount := 0
			deckLength := len(deckCopy)
			var tempCard c.Card

			for j := 0; j < deckLength; j++ {
				//If a card of the proper suit is found, add it to the count and store the last for removal
				if deckCopy[j].Suit == targetSuit {
					validCardCount++
					tempCard = deckCopy[j]
				}

			}

			//Calculate chance and append to list of individual probabilities
			tempFloat := float64(validCardCount) / float64(deckLength)
			indProbs = append(indProbs, tempFloat)
			//fmt.Print("[FLUSH TEST] valid/decksize: ", validCardCount, " / ", deckLength, " = ")
			//fmt.Printf("[FLUSH TEST] float: %f\n", tempFloat)

			//Remove the latest valid card to simulate drawing it
			var toRemove []c.Card
			toRemove = append(toRemove, tempCard)
			deckCopy = RemoveCards(deckCopy, toRemove)
		}

		//Multiply individual probabilities
		for i := 0; i < numSuitNeed; i++ {
			totalProb *= indProbs[i]
			//fmt.Println("[FLUSH TEST] probs: ", totalProb)
		}

		//Calculate number of orderings as the draw order doesn't matter (including free draws). Max is 7! = 5040
		numPermutations := Factorial(numSuitNeed)

		totalProb *= float64(numPermutations)

		return totalProb
	}

	//Find the chance of getting one card of a particular value (and suit if applicable), then store the chance for multiplication
	for i := 0; i < numNeeded; i++ {
		validCardCount := 0
		deckLength := len(deckCopy)
		var tempCard c.Card

		for j := 0; j < deckLength; j++ {
			//If the value is equal and the suit matches or isn't specified, add to count of valid cards
			if targetVals[i] == deckCopy[j].Val && (targetSuit == "" || targetSuit == deckCopy[j].Suit) {
				validCardCount++
				tempCard = deckCopy[j]
			}
		}

		//Calculate chance and append to list of individual probabilities
		tempFloat := float64(validCardCount) / float64(deckLength)
		indProbs = append(indProbs, tempFloat)
		//fmt.Print(validCardCount, " / ", deckLength, " = ")
		//fmt.Printf("%f\n", tempFloat)

		//Remove the latest valid card to simulate drawing it
		var toRemove []c.Card
		toRemove = append(toRemove, tempCard)
		deckCopy = RemoveCards(deckCopy, toRemove)
	}

	//fmt.Print("\n\n")

	//Multiply individual probabilities
	for i := 0; i < numNeeded; i++ {
		totalProb *= indProbs[i]
		//fmt.Printf("%f ", totalProb)
	}

	//Calculate number of orderings as the draw order doesn't matter (including free draws). Max is 7! = 5040
	numPermutations := Factorial(numNeeded)
	//fmt.Printf("\n%f ", numPermutations)

	totalProb *= float64(numPermutations)
	//fmt.Printf("%f", totalProb)

	return totalProb
}

// Determining probability of future hands from DetermineFutureHands
func DetermineFutureProbability(hand []c.Card, futureHands []string) []float64 {
	var futureProbs []float64
	canAddNumCards := 7 - len(hand)
	if Contains(futureHands, "One Pair") {
		pairVals := make([]int, 13)
		for _, card := range hand {
			pairVals[card.Val]++
		}
		firstPair := -1
		for index, count := range pairVals {
			if count >= 2 {
				firstPair = index
			}
		}
		//just multiply by num cards in hand and do
		onePairProb := 0.0
		if canAddNumCards == 0 {
			onePairProb = 0
		} else if firstPair != -1 {
			onePairProb = 1
		} else {
			for i := 1; i <= canAddNumCards; i++ {
				notOnePairProb := math.Pow(float64(52-4*len(hand))/float64(52-len(hand)), float64(canAddNumCards-i))
				onePairProb += float64(canAddNumCards) * notOnePairProb * math.Pow(float64(3*len(hand))/float64(52-len(hand)), float64(i))
			}
		}
		futureProbs = append(futureProbs, onePairProb)
	}
	if Contains(futureHands, "Two Pair") {
		pairVals := make([]int, 13)
		for _, card := range hand {
			pairVals[card.Val]++
		}
		firstPair := -1
		for index, count := range pairVals {
			if count >= 2 && firstPair == -1 {
				firstPair = index
			}
		}
		twoPairProb := 0.0
		if canAddNumCards == 0 {
			twoPairProb = 0
		} else if firstPair != -1 {
			for i := 1; i <= canAddNumCards; i++ {
				//basically 1 pair?
				notTwoPairProb := math.Pow(float64(52-4*len(hand))/float64(52-len(hand)), float64(canAddNumCards-i))
				twoPairProb += float64(canAddNumCards) * notTwoPairProb * math.Pow(float64(3*len(hand))/float64(52-len(hand)), float64(i))
			}
		} else if firstPair == -1 {
			for i := 1; i <= canAddNumCards; i++ {
				//1 card only or at least 1 card but no pairs
				notTwoPairProb := math.Pow(float64(52-4*len(hand))/float64(52-len(hand)), float64(canAddNumCards-i-2))
				otherPairProb := Factorial(26-2*len(hand)) / (Factorial(2) * Factorial(26-2*len(hand)-2)) * (float64(52-4*len(hand)) / 2) / float64(52-len(hand))
				twoPairProb += float64(canAddNumCards-2) * otherPairProb * notTwoPairProb * math.Pow(float64(3*len(hand))/float64(52-len(hand)), float64(i))
			}
		} else {
			//already 2 pairs in hand
			twoPairProb = 1.0
		}
		futureProbs = append(futureProbs, twoPairProb)
	}
	if Contains(futureHands, "Three of a Kind") {
		//either 1 pair, 2 pair, 3 pair, or just single cards
		pairVals := make([]int, 13)
		for _, card := range hand {
			pairVals[card.Val]++
		}
		firstPair := -1
		secondPair := -1
		thirdPair := -1
		alreadyThreeOfAKind := false
		currentProb := 0.0
		for index, count := range pairVals {
			if count == 3 {
				alreadyThreeOfAKind = true
				break
			}
			if count == 2 && firstPair == -1 && secondPair == -1 {
				firstPair = index
			}
			if count == 2 && firstPair != -1 && secondPair == -1 {
				secondPair = index
			}
			if count == 2 && firstPair != -1 && secondPair != -1 {
				thirdPair = index
			}
		}
		if canAddNumCards == 0 {
			currentProb = 0
		} else if alreadyThreeOfAKind {
			currentProb = 1.00
		} else {
			for i := 1; i < canAddNumCards; i++ {
				if thirdPair == -1 && secondPair == -1 && firstPair != -1 {
					//only 1 pair
					ThreeNotGottenProb := math.Pow(float64(1)/float64(52-len(hand)), float64(7-canAddNumCards+i))
					currentProb += Factorial(52-len(hand)) / (Factorial(canAddNumCards-i) * Factorial(52-len(hand)-canAddNumCards+i)) * ThreeNotGottenProb * math.Pow(float64(1)/float64(52-len(hand)), float64(canAddNumCards-i))
				} else if firstPair != -1 && secondPair != -1 && thirdPair == -1 {
					//2 pairs, half 1 pair prob
					currentProb += Factorial(52-len(hand)) / (Factorial(canAddNumCards-i) * Factorial(52-len(hand)-canAddNumCards+i)) * math.Pow(float64(1)/float64(52-len(hand)), float64(7-canAddNumCards+i)) * math.Pow(float64(1)/float64(52-len(hand)), float64(canAddNumCards-i)) / float64(2)
				} else if firstPair != -1 && secondPair != -1 && thirdPair != -1 {
					//3 pairs aka 1 card left, so can hardcode value, need 1 of 3 cards when 46 left
					currentProb = float64(3) / float64(46)
					break
				}
			}
		}
		futureProbs = append(futureProbs, currentProb)
	}
	if Contains(futureHands, "Full House") {
		tripleVals := make([]int, 13)
		for _, card := range hand {
			tripleVals[card.Val]++
		}
		firstTriple := -1
		secondTriple := -1
		firstPair := -1
		secondPair := -1
		thirdPair := -1
		currentProb := 0.00
		for index, count := range tripleVals {
			if count == 3 && firstTriple == -1 {
				firstTriple = index
			}
			if count == 3 && firstTriple != -1 && index != firstTriple {
				secondTriple = index
			}
			if count == 2 && firstPair == -1 && secondPair == -1 {
				firstPair = index
			}
			if count == 2 && firstPair != -1 && secondPair == -1 {
				secondPair = index
			}
			if count == 2 && firstPair != -1 && secondPair != -1 {
				thirdPair = index
			}
		}
		if firstPair != -1 && firstTriple != -1 {
			currentProb = 1.00
		} else {
			for i := 1; i < canAddNumCards; i++ {
				if firstPair == -1 && firstTriple == -1 {
					//only 1 pair
					ThreeNotGottenProb := math.Pow(float64(1)/float64(52-len(hand)), float64(7-canAddNumCards+i))
					currentProb += Factorial(52-len(hand)) / (Factorial(canAddNumCards-i) * Factorial(52-len(hand)-canAddNumCards+i)) * ThreeNotGottenProb * math.Pow(float64(1)/float64(52-len(hand)), float64(canAddNumCards-i))
				} else if firstPair != -1 && secondPair == -1 && firstTriple == -1 {
					//2 pairs, half 1 pair prob
					currentProb += Factorial(52-len(hand)) / (Factorial(canAddNumCards-i) * Factorial(52-len(hand)-canAddNumCards+i)) * math.Pow(float64(1)/float64(52-len(hand)), float64(7-canAddNumCards+i)) * math.Pow(float64(1)/float64(52-len(hand)), float64(canAddNumCards-i)) / float64(2)
				} else if secondPair != -1 && thirdPair == -1 && firstTriple == -1 {
					//3 pairs aka 1 card left, so can hardcode value, need 1 of 3 cards when 46 left
					currentProb = float64(3) / float64(46)
					break
				} else if thirdPair != -1 && firstTriple == -1 {
					//6 cards in hand, can't make full house
					currentProb = 0.00
					break
				} else if firstTriple != -1 && secondTriple != -1 {
					//1 card left, but a full house is already there
					currentProb = 1.00
					break
				}
			}
		}
		futureProbs = append(futureProbs, currentProb)
	}
	if Contains(futureHands, "Four of a Kind") {
		tripleVals := make([]int, 13)
		for _, card := range hand {
			tripleVals[card.Val]++
		}
		firstTriple := -1
		secondTriple := -1
		alreadyFourOfAKind := false
		for index, count := range tripleVals {
			if count == 4 {
				alreadyFourOfAKind = true
				break
			} else {
				if count == 3 && firstTriple == -1 {
					firstTriple = index
				}
				if count == 3 && firstTriple != -1 && index != firstTriple {
					secondTriple = index
				}
			}
		}
		fourOfAKindProb := 0.00
		if canAddNumCards == 0 {
			fourOfAKindProb = 0
		} else if alreadyFourOfAKind {
			fourOfAKindProb = 1.00
		} else {
			for i := 1; i <= canAddNumCards; i++ {
				if secondTriple == -1 && firstTriple != -1 {
					if canAddNumCards != 0 {
						//only 1 triple so it's easier
						FourNotGottenProb := math.Pow(float64(48)/float64(52-len(hand)), float64(7-canAddNumCards+i))
						fourOfAKindProb += Factorial(52-len(hand)) / (Factorial(canAddNumCards-i) * Factorial(52-len(hand)-canAddNumCards+i)) * FourNotGottenProb * math.Pow(float64(1)/float64(52-len(hand)), float64(canAddNumCards))
					} else {
						fourOfAKindProb = float64(0)
						break
					}
				} else if secondTriple != -1 && firstTriple != -1 {
					if canAddNumCards != 0 {
						//should only be 6 cards so just double the prob of grabbing 1 card? (46 cards left, need 2 specific ones)
						fourOfAKindProb = float64(1) / float64(23)
						break
					} else {
						fourOfAKindProb = float64(0)
						break
					}
				}
			}
		}
		futureProbs = append(futureProbs, fourOfAKindProb)
	}
	return futureProbs
}

// Determining what hands can be created using the current hand and cards in the deck
func DetermineFutureHands(hand []c.Card, currentHands []string) []string {
	futureHands := make([]string, 0)
	if len(hand) < 7 {
		if Contains(currentHands, "Full House") {
			futureHands = append(futureHands, "Placeholder - Full House")
		} else if Contains(currentHands, "Three of a Kind") {
			futureHands = append(futureHands, "Four of a Kind")
		} else if Contains(currentHands, "One Pair") {
			futureHands = append(futureHands, "Three of a Kind")
			if len(hand) < 6 {
				futureHands = append(futureHands, "Four of a Kind")
				futureHands = append(futureHands, "Two Pair")
			}
		} else if Contains(currentHands, "Two Pair") {
			futureHands = append(futureHands, "Three of a Kind")
			if len(hand) < 6 {
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

func GetHandArray(hand *h.Hand) []c.Card {
	handArray := make([]c.Card, 0)
	handArray = append(handArray, hand.ActualHand...)
	return handArray
}

// Checking the hand type (i.e. one pair, 3 of a kind...) of the hand at the current time
func CheckHandType(hand []c.Card) []string {
	doesHandHaveType := false
	handTypes := make([]string, 0) //store hand types
	cardCount := make(map[int]int) //store count of card numbers
	pairCount := 0
	trioCount := 0
	quartetCount := 0
	for _, card := range hand {
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

// Removes desired cards from an array of cards
func RemoveCardsFromArray(cards []c.Card, extraCards []c.Card) (cardCopy []c.Card) {
	cardCount := len(cards)
	copyCount := len(extraCards)

	//Check each card of the deck against input. Do not add to deckCopy if card exists in hand or community cards.
	for j := 0; j < cardCount; j++ {
		addCheck := true

		for i := 0; i < copyCount; i++ {
			if extraCards[i].Val == cards[j].Val && cards[i].Suit == cards[j].Suit {
				addCheck = false
			}
		}

		if addCheck {
			cardCopy = append(cardCopy, cards[j])
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

// Check the probability of a high card
func HighCard(deck Deck, cards []c.Card) float64 {
	if len(cards) > 0 {
		return 1.00
	} else {
		return 0.00
	}
}

// Check whether the hand and community cards provided can form a royal straight flush with the rest of the deck and return boolean and probability
func RoyalFlush(deck Deck, cards []c.Card) float64 {
	cardCount := len(cards)
	remaining := 7 - cardCount

	//fmt.Println("[FRONTEND CHECK] count: ", cardCount)

	if cardCount == 0 {
		return 0.00
	}

	//Variables for probability calculation (which is commented out)
	//deckCopy := RemoveCards(deck, cards)
	//deckCount := len(deckCopy)
	//chosenSuit := ""

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

	suits := map[string]struct{}{}
	royalFlush := make(map[int]c.Card)
	needCards := [5]bool{true, true, true, true, true}
	var currVals []int
	targetSuit := ""
	hR := 0 //Heart-Royal
	dR := 0 //Diamond-Royal
	cR := 0 //Club-Royal
	sR := 0 //Spade-Royal

	var dupe c.Card
	var tempCard c.Card

	needCount := 5

	//Populate royalFlush map with cards of like suit for values of 12, 11, 10, 9, 0
	for i := 0; i < cardCount; i++ {
		tempCard = cards[i]

		//Count suits
		suits[tempCard.Suit] = struct{}{}

		//Copy down values
		currVals = append(currVals, tempCard.Val)

		//If there are 4 different suits, then 5/7 cards cannot be the same suit for a flush
		if len(suits) == 4 {
			return 0.00
		}

		if ContainsInt([]int{12, 11, 10, 9, 0}, tempCard.Val) {
			switch tempCard.Suit {
			case "Heart":
				hR++
			case "Diamond":
				dR++
			case "Club":
				cR++
			case "Spade":
				sR++
			}
		}

		//Switch case to ignore all cards not needed for a royal flush straight
		switch tempCard.Val {
		case 12, 11, 10, 9, 0:

			//If map slot is empty, add card
			if royalFlush[tempCard.Val].Suit == "" {
				//fmt.Println("[FRONTEND CHECK] map initialization check: success")
				royalFlush[tempCard.Val] = tempCard
				//chosenSuit = tempCard.Suit

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

		skipLoop := false

		//check if card is found from the 5 needed
		if !needCards[i] {
			for j := 0; j < 5; j++ {

				//fmt.Println("i: ", i, "=", needCards[i], " j: ", j, "=", needCards[j])

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
						secondIndex = 12
					case 1:
						secondIndex = 11
					case 2:
						secondIndex = 10
					case 3:
						secondIndex = 9
					case 4:
						secondIndex = 0
					}

					//fmt.Println("Comparing: ", c.GetCardName(royalFlush[firstIndex]), " and ", c.GetCardName(royalFlush[secondIndex]))

					//compare card suits and remove any royal cards that don't belong to the primary suit
					if royalFlush[firstIndex].Suit != royalFlush[secondIndex].Suit && royalFlush[secondIndex].Suit != "" {
						//fmt.Println("Main loop suit inconsistency")

						compare1 := 0
						compare2 := 0

						switch royalFlush[firstIndex].Suit {
						case "Heart":
							compare1 = hR
						case "Diamond":
							compare1 = dR
						case "Club":
							compare1 = cR
						case "Spade":
							compare1 = sR
						}

						switch royalFlush[secondIndex].Suit {
						case "Heart":
							compare2 = hR
						case "Diamond":
							compare2 = dR
						case "Club":
							compare2 = cR
						case "Spade":
							compare2 = sR
						}

						if compare1 > compare2 {
							//fmt.Println("Deleting second comparison: ", c.GetCardName(royalFlush[secondIndex]))
							switch royalFlush[secondIndex].Val {
							case 12:
								needCards[0] = true
							case 11:
								needCards[1] = true
							case 10:
								needCards[2] = true
							case 9:
								needCards[3] = true
							case 0:
								needCards[4] = true

							}

							delete(royalFlush, secondIndex)

							needCount++

							skipLoop = true

						} else {
							//fmt.Println("Deleting first comparison: ", c.GetCardName(royalFlush[firstIndex]))
							switch royalFlush[firstIndex].Val {
							case 12:
								needCards[0] = true
							case 11:
								needCards[1] = true
							case 10:
								needCards[2] = true
							case 9:
								needCards[3] = true
							case 0:
								needCards[4] = true

							}

							delete(royalFlush, firstIndex)

							needCount++

							skipLoop = true

						}

					} else {
						targetSuit = royalFlush[firstIndex].Suit
					}
				}

				if skipLoop {
					break
				}
			}
		}
	}

	//fmt.Println("[FRONTEND CHECK] needcount: ", needCount, " remaining: ", remaining)

	//If more cards are needed than should be drawn, return false and 0.0%
	if needCount > remaining {
		//fmt.Println("First case (impossible)")
		return 0.00
	} else if needCount == 0 {
		//fmt.Println("Second case (found)")
		return 1.00 //if none needed, return true and 100.0%
	}

	//Find how many royal flushes are possible and calculate their probabilities
	if cardCount <= 2 {
		//fmt.Println("[FRONTEND CHECK] Secondary loop all available")
		//All royal flushes available
		targetSuit = ""
	} else {
		//fmt.Println("[FRONTEND CHECK] Secondary loop suits check")
		prob := 0.00

		for suitIndex := 0; suitIndex < 4; suitIndex++ {
			//Assign amount needed based on suit
			currSuit := 5
			switch suitIndex {
			case 0:
				currSuit -= hR
				targetSuit = "Heart"
			case 1:
				currSuit -= dR
				targetSuit = "Diamond"
			case 2:
				currSuit -= cR
				targetSuit = "Club"
			case 3:
				currSuit -= sR
				targetSuit = "Spade"
			}

			//Check if possible for the suit and then assign targetVals appropriately
			if currSuit <= remaining {
				//fmt.Println("[FRONTEND CHECK] Tertiary loop individual suits check. Current: ", targetSuit)
				//Define targetVals
				var targetVals []int

				i := 12
				for len(targetVals) < needCount {
					//If value is NOT contained (to avoid searching for duplicates) add to targetVals. 5-8 were included to allow for value overlap between multiple suits (calc is the same 1/deck size)
					if !ContainsInt([]int{12, 11, 10, 9, 8, 7, 6, 5, 0}, i) {
						targetVals = append(targetVals, i)
					}
					i--
				}

				//fmt.Println("[FRONTEND CHECK] Tertiary loop prob before: ", prob)

				//Add probabilities for each hand together
				prob += FindCardProb(cards, targetVals, targetSuit, 0)

				//fmt.Println("[FRONTEND CHECK] Tertiary loop prob after: ", prob)
			}
		}

		//fmt.Println("Third case (some)")
		return prob
	}

	//Define targetVals
	var targetVals []int

	for i := 0; i < 5; i++ {
		if needCards[i] {
			switch i {
			case 0:
				targetVals = append(targetVals, 12)
			case 1:
				targetVals = append(targetVals, 11)
			case 2:
				targetVals = append(targetVals, 10)
			case 3:
				targetVals = append(targetVals, 9)
			case 4:
				targetVals = append(targetVals, 0)
			}
		}
	}

	prob := 0.00

	for i := 0; i < 4; i++ {
		currSuit := 5
		switch i {
		case 0:
			currSuit -= hR
			targetSuit = "Heart"
		case 1:
			currSuit -= dR
			targetSuit = "Diamond"
		case 2:
			currSuit -= cR
			targetSuit = "Club"
		case 3:
			currSuit -= sR
			targetSuit = "Spade"
		}

		if currSuit == 5 {
			baseTarget := []int{12, 11, 10, 9, 0}
			prob += FindCardProb(cards, baseTarget, targetSuit, 0)
			//fmt.Printf("[FRONTEND CHECK] +base prob: %f\n", prob)
		} else {
			//Find true probability
			prob += FindCardProb(cards, targetVals, targetSuit, 0)
			//fmt.Printf("[FRONTEND CHECK] +spec prob: %f\n", prob)
		}
	}
	//fmt.Println("Fourth case (all)")
	return prob
}

// Returns probability of a straight with a boolean for if it is a straightFlush.
func StraightCheck(deck Deck, cards []c.Card) (float64, float64) {
	//deckCopy := RemoveCards(deck, cards)
	cardCount := len(cards)

	if cardCount == 0 {
		return 0.00, 0.00
	}

	cards = ValSortCardsAsc(cards)
	flushBool := false
	lowVal := 0
	lowestVal := 0
	highVal := 0
	highestVal := 0
	chain := 0
	var currVals []int
	currSuit := ""

	//Checks the cards in hand for a straight by looking for 5 sequentially valued cards
	for i := 0; i < cardCount-1; i++ {
		tempCard := cards[i]
		tempNext := cards[i+1]
		diff := tempNext.Val - tempCard.Val

		if i == 0 {
			currVals = append(currVals, tempCard.Val)
			lowestVal = tempCard.Val
			highestVal = tempCard.Val
			currSuit = tempCard.Suit
		}
		if cardCount != 1 {
			currVals = append(currVals, tempNext.Val)
			if tempNext.Val < lowestVal {
				lowestVal = tempNext.Val
			}
			if tempNext.Val > highestVal {
				highestVal = tempNext.Val
			}
			if currSuit != "" && tempNext.Suit != currSuit {
				currSuit = ""
			}
		}

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

	//Loop is skipped if there is only 1 card
	if cardCount == 1 {
		lowVal = cards[0].Val
		lowestVal = cards[0].Val
		highVal = cards[0].Val
		highestVal = cards[0].Val
		currVals = append(currVals, cards[0].Val)
	}

	broadwayBool := (lowVal == 9 && highVal == 12 && chain == 4 && cards[0].Val == 0)

	//Check for extra cards
	var extraCards []c.Card
	var straightCards []c.Card

	for i := 0; i < cardCount; i++ {
		if cards[i].Val < lowVal && !(broadwayBool && cards[i].Val == 0) {
			extraCards = append(extraCards, cards[i])
		} else if cards[i].Val > highVal && !(broadwayBool && cards[i].Val == 0) {
			extraCards = append(extraCards, cards[i])
		}
	}

	//If there are extra cards, remove them and perform a FlushCheck
	if len(extraCards) > 0 {
		straightCards = RemoveCardsFromArray(cards, extraCards)
	} else {
		straightCards = cards
	}

	targetSuit := currSuit
	permaWrong := false

	//Run flush check to test for a straight flush
	for i := 0; i < len(straightCards); i++ {
		//fmt.Println("Straight Card #", i, ": ", c.GetCardName(straightCards[i]))
		if targetSuit == "" {
			targetSuit = straightCards[0].Suit
		} else if straightCards[i].Suit != targetSuit {
			flushBool = false
			permaWrong = true
		} else if !permaWrong {
			flushBool = true
		}
	}

	if flushBool {
		targetSuit = straightCards[0].Suit
	}

	//BROADWAY STRAIGHT (functionally the same as a straight with ace high but not flush)
	if broadwayBool {
		if flushBool {
			//fmt.Println("Broadway flush??")
			return 1.00, 1.00
		} else {
			return 1.00, 0.00
		}
	}

	//Regular straight
	if highVal-lowVal == 4 && chain == 5 {
		if flushBool {
			//fmt.Println("Normal straight flush??")
			return 1.00, 1.00
		} else {
			//return 1.00, 0.00
		}
	}

	//Begin 0 < x < 1 calcs
	prob := 0.00
	straightFlushProb := 0.00
	numToDraw := 7 - cardCount

	if numToDraw == 0 {
		return prob, straightFlushProb
	}

	rangeVal := 0
	highRangeVal := 0
	targetSizeOffset := 0
	var targetVals []int
	//numSeq := 0

	//fmt.Println("lowestVal: ", lowestVal)

	//Establish lowest end of range
	if cardCount <= 2 || lowestVal < numToDraw {
		rangeVal = 0
	} else {
		rangeVal = lowestVal - numToDraw
	}

	//Establish highest end of range
	if cardCount <= 2 || highestVal > (12-numToDraw) {
		highRangeVal = 12
	} else {
		highRangeVal = highestVal + numToDraw
	}

	//Grab 5 values in sequential order including current values and repeat until all groupings of 5 have been found
	for rangeVal <= (highRangeVal) {
		if !ContainsInt(currVals, rangeVal) {
			targetVals = append(targetVals, rangeVal)
		} else {
			//Simulate adding current values to targetVals
			targetSizeOffset++
		}

		//fmt.Println("rangeVal: ", rangeVal, " length: ", len(targetVals), " offset: ", targetSizeOffset)

		if len(targetVals)+targetSizeOffset == 5 {
			//if numSeq == 0 {

			if !broadwayBool && !(highVal-lowVal == 4 && chain == 5) {
				prob += FindCardProb(cards, targetVals, "", 0) //change "" to currSuit to begin implementing straight flush prob calc
			} else {
				prob = 1.00
			}

			if flushBool {
				straightFlushProb += FindCardProb(cards, targetVals, targetSuit, 0)
			}
			//fmt.Printf("Prob: %f\n", prob)
			//}
			//numSeq++
			targetVals = []int{}
			rangeVal -= 3 //Go back 3 vals to check the next sequence from 1 higher
			targetSizeOffset = 0
		} else {
			rangeVal++
		}
	}

	return prob, straightFlushProb
}

func FlushCheck(deck Deck, cards []c.Card) float64 {
	//deckCopy := RemoveCards(deck, cards)
	cardCount := len(cards)

	if cardCount == 0 {
		return 0.00
	}

	suits := make(map[string]int)
	var tempCard c.Card
	hR := 0
	dR := 0
	cR := 0
	sR := 0

	//Populate suits map with count of each suit
	for i := 0; i < cardCount; i++ {
		tempCard = cards[i]
		switch tempCard.Suit {
		case "Heart":
			hR++
		case "Diamond":
			dR++
		case "Club":
			cR++
		case "Spade":
			sR++
		}

		if suits[tempCard.Suit] > 0 {
			suits[tempCard.Suit]++
		} else {
			suits[tempCard.Suit] = 1
		}

		//If there are 4 different suits, then 5/7 cards cannot be the same suit for a flush
		if len(suits) == 4 {
			return 0.00
		}
	}

	//fmt.Println("[FLUSH CHECK] hr:", hR, " dr:", dR, " cr:", cR, " sR:", sR)

	if hR >= 5 || dR >= 5 || cR >= 5 || sR >= 5 {
		return 1.00
	}

	//proceed with calculations
	remaining := 7 - cardCount
	targetSuit := ""

	prob := 0.00

	for suitIndex := 0; suitIndex < 4; suitIndex++ {
		//Assign amount needed based on suit
		currSuit := 5
		switch suitIndex {
		case 0:
			currSuit -= hR
			targetSuit = "Heart"
		case 1:
			currSuit -= dR
			targetSuit = "Diamond"
		case 2:
			currSuit -= cR
			targetSuit = "Club"
		case 3:
			currSuit -= sR
			targetSuit = "Spade"
		}

		//Check if possible for the suit and then find probability
		if currSuit <= remaining {
			prob += FindCardProb(cards, []int{}, targetSuit, currSuit)
			//fmt.Println("[FLUSH TEST] total prob: ", prob)
		}
	}

	return prob
}

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
		fmt.Println("Invalid range (int 0-12, suit Heart/Diamond/Club/Spade) [PART OF TestGetCardIndex4]")
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
