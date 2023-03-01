Sprint 2 Progress:

  Frontend:
  
  - Completed end-to-end integration, allowing information to be sent between the frontend and backend.
  - 
  
  Backend:
  
  - Completed end-to-end integration, allowing information to be sent between the frontend and backend.
  - Completed check functions for most hand types:
    - Pairs
    - Two-pair
    - Three of a kind
    - Four of a kind
    - Full house
    - Straight
    - Royal flush

Frontend Unit Tests:

  - 

Backend Unit Tests (Golang tests):

  - Deck:
    - 5 GetCardIndex tests for accurate index returns from a deck.
    - GetCardName test for accurate proper card name return.
    - PrintDeck test for simply viewing full deck printout (no intended function yet, may be used to show cards remaining in deck).
    - RemoveCards test for accurate index offsets when 3 cards are removed from the output card array.
    - RoyalFlush test for identifying if an array of cards contains a royal flush (substitution with Flush and Straight functions planned).
    - 2 StraightCheck tests for identifying if an array of cards contains either a regular straight or a royal straight.
    - TestOnePairCheck test for checking if an array of cards can be identified as having a one pair (2 cards of the same number but not same suit) within it
    - TestTwoPairCheck test for checking if an array of cards can be identified as having a two pair (2 pairs, each 2 cards of the same number but not same suit) within it
    - TestThreeFourFullCheck test for checking if an array of cards can be identified as having 3 of a kind (3 cards of the same number but not same suit), 4 of a kind (all 4 cards of the same number but not same suit), or full house (one 1 pair and one 3 of a kind within a hand) within it. These tests were performed within the same testing function for easier comprehension.
    
  - Hand:
    - TestHandCardValues test for checking if card values and suits successfully insert into an array of card objects
    - TestHandAddOverflow test to check if adding more than 7 cards gives an error
  
Backend API Documentation:

  - 
