Sprint 2 Progress:

  Frontend:
  
  - Completed end-to-end integration, allowing information to be sent between the frontend and backend.
  - Created assets for the display of card values according to backend data.
  - Designed wireframes for and implamented the basic structure of the page's layout.
  
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

  - Basic test to check if the app is properly created
  - Checks if the number of image locations stored in currImgs is equal to the amount of cards in currHand.
  - Ensures that the number of elements rendered in the right side bar is equal to the number of elements in currHand.
  - Ensures that the number of matLabel elements rendered is equal to the number of entries in currImgs.
  - Compares the content of currHand to the values displayed to ensure that they are equal.
  - Compares the content of currImg to the src attribute of the images displayed to ensure that they are roughly equal.
  - Compares the src attribute of the displayed images to the content of currHand to see if the proper image is displayed with respect to the stored hand. 

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

  - API url: localHost:4200
  - POST
    - The URL parameters are /api/hand.
    - The parameters that are sent in the POST request includes a struct with an Index as an int, Suit as a string, and Val as an int.
    - From looking at the network tab when inspecting elements on chrome, a sample POST request looks like: {"Suit":"club","Val":1,"Index":0}
    - On a successful POST, the network tab displayed that there was no response data to show. 
  - GET
    - The URL parameters are /api/hand.
    - The data parameters for this GET request is an array containing structs with an Index as an int, Suit as a string, and Val as an int.
    - The network tab shows that a GET request immediately the POST request above has the body: [{Val: 1, Suit: "club", Index: 0}].
    - The success response shows the same array: [{"Val":1,"Suit":"club","Index":0}] but now its strored as variable on the front end.
