Sprint 3 Progress:

Frontend:

  -g

Backend:

  - Added Probabilitys and HandProb structure to return float64s for each handtype to the frontend and integrated it with the API.
  - Adjusted probability functions to be compatible with the new structure.
  - Linked probability functions to the new structure to enable handtype identification on the frontend. Localhost:5000/api/prob shows each chance. (100% or 0%)
  - Added FindCardProb to take in targeted card values and a suit to return the probability of drawing each of the targeted values in any order.
  - Began integration of FindCardProb into some of the handtype identifier functions to start returning more than just certain (100% or 0%) probabilities.
  - Added RemovecardsFromArray. Similar functionality to RemoveCards except with solely arrays and not a deck object.
  - Added ContainsInt function for ints in addition to the preexisting one, Contains, for strings.
  - Added DetermineFutureProbability function, in progress, but should determine the probability to draw a given hand type in future draws. For now One Pair and partially Three of a Kind should work.

Frontend Unit Tests:

  -g

Backend Unit Tests:

  - Deck:
    - TestRemoveCardsFromArray to test proper remove of cards from an array of cards. Functionally identical to RemoveCards function for deck objects.
    - TestFindCardProb tests a variety of hypothetical statistics situations for a variety of handtypes. Test0 compares 5-draw royal flush chance to Wikipedia's.
    - 5 GetCardIndex tests for accurate index returns from a deck.
    - GetCardName test for accurate proper card name return.
    - PrintDeck test for simply viewing full deck printout (no intended function yet, may be used to show cards remaining in deck).
    - RemoveCards test for accurate index offsets when 3 cards are removed from the output card array.
    - RoyalFlush test for identifying if an array of cards contains a royal flush (substitution with Flush and Straight functions planned).
    - 2 StraightCheck tests for identifying if an array of cards contains either a regular straight or a royal straight.
    - TestOnePairCheck test for checking if an array of cards can be identified as having a one pair (2 cards of the same number but not same suit) within it
    - TestTwoPairCheck test for checking if an array of cards can be identified as having a two pair (2 pairs, each 2 cards of the same number but not same suit) within it.
    - TestThreeFourFullCheck test for checking if an array of cards can be identified as having 3 of a kind (3 cards of the same number but not same suit), 4 of a kind (all 4 cards of the same number but not same suit), or full house (one 1 pair and one 3 of a kind within a hand) within it. These tests were performed within the same testing function for easier comprehension. 
    - TestFutureProbabilityFourOfKind test for checking if the future probability correctly returns for a hand containing only one three of a kind and nothing else.
  
  - Hand:
    - TestHandCardValues test for checking if card values and suits successfully insert into an array of card objects
    - TestHandAddOverflow test to check if adding more than 7 cards gives an error

Backend API Documentation:

  - API url: localHost:4200
  - POST
    - The URL parameters are /api/hand.
    - The parameters that are sent in the POST request includes a struct with an Index as an int, Suit as a string, and Val as an int.
    - From looking at the network tab when inspecting elements on chrome, a sample POST request looks like: {"Suit":"club","Val":1,"Index":0}
    - On a successful POST, the network tab displayed that there was no response data to show. 
  - GET
    - The URL parameters are /api/hand and /api/prob.
    - The data parameters for this GET request is an array containing structs with an Index as an int, Suit as a string, and Val as an int.
    - The network tab shows that a GET request immediately the POST request above has the body: [{Val: 1, Suit: "club", Index: 0}].
    - The success response shows the same array: [{"Val":1,"Suit":"club","Index":0}] but now its strored as variable on the front end.
    - /api/prob shows the array of HandProb structs run through UpdateProb for each handtype. Results are currently limited to 1.00 for an identified handtype, and 0.00 if one was not found.
