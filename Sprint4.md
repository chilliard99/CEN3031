Sprint 4 Progress:

Frontend:

  - Added content for the Hand List page
  - Filled in Poker Terms page
  - Added Probabilities on simulation page
  - Added darkmode for the simulation page
  - Various UI Changes (Hover animations, positioning improvements)


Backend:

  - Enhanced RoyalFlush with probability calculations to give responses between 0.00 and 1.00.
  - Enhanced StraightCheck with probability calculations to give responses between 0.00 and 1.00.
  - Created and enhanced FlushCheck with probability calculations to give responses between 0.00 and 1.00.
  - Troubleshooted a wide variety of minor problems that we were unable to test for until this stage of development.
  - Cleaned comments

Frontend Unit Tests:

  - Test to ensuare that darkmode elements are not displayed when darkmode is not toggled
  - Test to ensure that darkmode elements are displayed when darkmode is toggled
  - Test to ensure that the function toggleDarkmode toggles darkmode. 
  - Test to check if currentProb has a member for each hand type
  - Test to check if the sum of current probs is 0 at the start of the program.
  - Test to check if testForRepeats can detect a repeat face value card
  - Test to ensure that changeTab function switches the matTab
  - Test to check that addCard is not called if a repeat face card is passed into setVal.

Backend Unit Tests:

  - Deck:
    - TestRemoveCardsFromArray to test proper remove of cards from an array of cards. Functionally identical to RemoveCards function for deck objects.
    - TestFindCardProb tests a variety of hypothetical statistics situations for a variety of handtypes. Test0 compares 5-draw royal flush chance to Wikipedia's.
    - 5 GetCardIndex tests for accurate index returns from a deck.
    - GetCardName test for accurate proper card name return.
    - PrintDeck test for simply viewing full deck printout (no intended function yet, may be used to show cards remaining in deck).
    - RemoveCards test for accurate index offsets when 3 cards are removed from the output card array.
    - RoyalFlush test for identifying if an array of cards contains a royal flush (NEW added subtests for 0 < x < 1 output).
    - StraightCheck tests for identifying if an array of cards contains either a regular straight or a straight flush (NEW added substests for 0 < x < 1 output).
    - FlushCheck tests for identifying if an array of cards contains a flush (NEW added substests for 0 < x < 1 output).
    - DebugLogic helps test for issues with RoyalFlush and CheckStraight due to the similarities in the calculation programming.
    - TestFactorial ensures that the factorial function produces proper results.
    - 2 tests for the card sorting functions to ensure they are sorting correctly.
    - TestMassTest simply runs 50 5-card draws and prints out the cards and relevant probability calculation results to look for outstanding errors. Manual test requires a fail statement to view outputs.
    - TestOnePairCheck test for checking if an array of cards can be identified as having a one pair (2 cards of the same number but not same suit) within it
    - TestTwoPairCheck test for checking if an array of cards can be identified as having a two pair (2 pairs, each 2 cards of the same number but not same suit) within it.
    - TestThreeFourFullCheck test for checking if an array of cards can be identified as having 3 of a kind (3 cards of the same number but not same suit), 4 of a kind (all 4 cards of the same number but not same suit), or full house (one 1 pair and one 3 of a kind within a hand) within it. These tests were performed within the same testing function for easier comprehension. 
    - TestFutureProbabilityFourOfKind test for checking if the future probability correctly returns for a hand containing only one three of a kind and nothing else.
    - Altered TestFutureHand test to accommodate GetHandArray function.
  
  - Hand:
    - TestHandCardValues test for checking if card values and suits successfully insert into an array of card objects
    - TestHandAddOverflow test to check if adding more than 7 cards gives an error
    - testRandomAndRemove e2e cypress test to test front end funcionality of random hand and remove functions.
    - testRoyal e2e cypress test to show user manually inputing a hand that contains a royal flush.

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
    - /api/prob shows the array of HandProb structs run through UpdateProb for each handtype.
