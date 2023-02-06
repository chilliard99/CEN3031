User Stories:

     front:
  
     back:
     - As a back-end engineer, I want the most basic testing environment in Go so that we can begin testing our back-end implementations.
     - As a developer, I would like card and deck variables in the back-end so that we can begin implementing functions to assist users via future functions.
     - As a new user, I would like to be able to learn the names of the cards so that I can identify them without seeing them.
  
Issues:

     front:

     back:
     - Back-end Testing: Starting from scratch meant no testing structure, and so a basic project and data structure was needed to begin testing the code through Go.
     - Card deck structure: A more complex data structure for storing cards in a deck with constructor and accessors functions would allow us to store a full deck of cards and sort through them or shuffle as you would with a real deck.

Completed:

     front:

     back:
     - Back-end Testing: A simple Deck struct and file was created with a function to return a hard-coded deck size. This was run through a Go test function with the appropriate passes and fails.
     - Card deck structure: Card and Deck structs were made with a deck constructor that would fill the deck with 52 cards, 13 for each of the 4 suits with no duplicates. Additionally an accessor function would return the index of a card in the deck array when provided with a card value int and a suit string. The test for the accessor function passed and failed when paired with correct and incorrect inputs respectively.

Not-Completed:

     front: 
     - Light/Dark Mode Toggle: This was not completed since we are still building out the the user interface and we will add this feature later in the development process.

     back:
