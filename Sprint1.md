Code Demonstration:

     Front:
     - https://youtu.be/riwa8QpBbl0
     
     Back:
     - https://youtu.be/zJYw1-95shE

User Stories:

     front:
     - As a front end dev, i want to use angular material which makes designing components easier, so we can speed up development and add more user features.
     - As a fronted developer I want to use Angular Material to display some card images on the page
     - As a frontend dev, I want users to be able to select cards that they wish to add to a given hand so that the program holds some basic functionality.
     - As a user I would like to be able to toggle light and dark themes
     - As a frontend developer, I want to have a basic display so that we can ensure that Angular is set up correctly and further expand functionality.
  
     back:
     - As a back-end engineer, I want the most basic testing environment in Go so that we can begin testing our back-end implementations.
     - As a developer, I would like card and deck variables in the back-end so that we can begin implementing functions to assist users via future functions.
     - As a new user, I would like to be able to learn the names of the cards so that I can identify them without seeing them.
     - As a backend dev, I'd like to store past hands users have selected so that the program stores more information that the user can call from.
     - As a developer, I'd like to write unit tests for each data structure so that we can see if they are working as intended.
     - As a user, I'd like to be able to have an account on the program so I can see only my own past hands/cards.
     - As a user, I'd like to be able to see the past 5 submitted hands on the homepage of the program so it looks more appealing.
  
Issues:

     front:
     - Displaying Basic Cards: Have a basic card displayed through an Angular material component.
     - Card Selection feature: Allow users to a suit and value for a card and display that card.
     - Add Angular Material
     - Implament Basic Back End Connection: Have some sort of communication between back and front end. 
     - Light/Dark Mode Toggle: Allow the user to toggle between a light and dark mode on the page.
     - Complete Front/Backend Connection: Allow complete transfer of card information between the front and back end. 

     back:
     - Back-end Testing: Starting from scratch meant no testing structure, and so a basic project and data structure was needed to begin testing the code through Go.
     - Card deck structure: A more complex data structure for storing cards in a deck with constructor and accessors functions would allow us to store a full deck of cards and sort through them or shuffle as you would with a real deck.
     - Card names: To make the data more user-friendly and teach users more about the game, a function that inputs a card and returns its proper name would be good.
     - Hand data structure: A data structure separate from the deck that will be used to draw cards from the deck and determine what type of hand a player has at any given time. The latter function has yet to be implemented as of this sprint.
     - Storing past 5 hands for the front end: This should be the next issue worked on as it is a natural extension of having a working Hand data structure, although the front end and back end will need to have a working connection for the results to be displayable on a webpage.
     - Unit tests for data structures: This issue is still being worked on as the Deck and Hand data structures may still need to be expanded in the future.
     - User Accounts to see past hands: This will require a connection between the front and back ends, so this issue can't be started on until this happens first.

Completed:

     front:
     - Add Angular Material: Angular material was added as a dependency for the project and we were able to display material components.
     - Implament Basic Back End Connection: A user can now use the two inputs and button to add to a hand in the back end.
     - Display a Basic Card: We created a card and can display it on our page. 

     back:
     - Back-end Testing: A simple Deck struct and file was created with a function to return a hard-coded deck size. This was run through a Go test function with the appropriate passes and fails.
     - Card deck structure: Card and Deck structs were made with a deck constructor that would fill the deck with 52 cards, 13 for each of the 4 suits with no duplicates. Additionally an accessor function would return the index of a card in the deck array when provided with a card value int and a suit string. The test for the accessor function passed and failed when paired with correct and incorrect inputs respectively.
     - Card names: A function was created to take the int value of a card and convert it into written word format as a string and concatenate it with the suit (i.e. "King of Clubs").
     - Hand data structure: Basic hand functionality (drawing a card with a specific suit and value) has been implemented, and tests involving it have passed.

Not-Completed:

     front: 
     - Light/Dark Mode Toggle: This was not completed since we are sill building out the the user interface and we will add this feature later in the developmt process.
     - Complete Front end Back end connections: This one is in progress, as cody added the basic connection that we can build off of inorder to display data from the backend.
     - Card Selection Feature: This feature is only partially completed, a user can input a suit and value, but it does not interact with other elements in the frontend nor the backend. 

     back:
     - Storing past 5 hands for the front end: Not completed because this requires a separate data structure for storing all the hands as well as a connection between the front and back ends to display the past 5 hands.
     - Unit tests for data structures: This issue is still being worked on as the Deck and Hand data structures may still need to be expanded in the future, as well as adding unit tests for any data structures we will create in the future.
     - User Accounts to see past hands: This will require a connection between the front and back ends, so this issue can't be started on until this happens first.