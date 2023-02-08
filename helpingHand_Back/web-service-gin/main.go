package main

import (
	"example/web-service-gin/deck"
	"example/web-service-gin/hand"
	"example/web-service-gin/handler"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	//Menu implemented to choose between running the program and varied backend testing
	menuOn := true

	fmt.Println("Welcome to Helping Hand's backend environment")
	fmt.Println("To start the program, press enter")
	fmt.Println("Options:\n0: Exit\n1: Generate a new deck\n2: Add card to hand\n3: Find card index in deck\n4: Get card name from hand\n5: Reprint menu")

	for menuOn {

		var menuInput string
		var currDeck deck.Deck
		currentHand := hand.NewHand("initial")

		fmt.Scanln(&menuInput)

		switch menuInput {
		case "":
			//Start the program
			//All of the previous main items:
			currDeck = deck.NewDeck()
			fmt.Println(deck.GetCardIndex(currDeck, 0, "Heart"))
			currentHand = hand.NewHand("initial")
			hand.AddCardHand(currentHand)
			fmt.Println("current length of hand is: ")
			fmt.Println(len(currentHand.ActualHand))

			r := gin.Default()
			api := r.Group("/api")
			{
				api.GET("/ping", handler.PingGet())
				api.GET("/hand", handler.HandGet(currentHand))
				api.POST("/hand", handler.HandPost(currentHand))
			}
			r.Run("0.0.0.0:5000")

		case "0":
			menuOn = false

		case "1":
			//broaden deck options?
			/*
			//Deck options
			fmt.Scanln(&menuInput)

			switch menuInput {
			*/

			
			//Generate a new deck
			currDeck = deck.NewDeck()
			}
		case "2":
			//Add a card to hand
			hand.AddCardHand(currentHand)

		case "3":
			//Find card index in deck

			var valueInput int
			var suitInput string

			//Input handler for card value
			valid := true

			for valid {
				fmt.Println("Enter a card value (e.g. 0, 1, 2,..., 12): ")
				fmt.Scanf("%d", &valueInput)

				//input := bufio.NewScanner(os.Stdin)
				//input.Scan()
				//valueInput: = strconv.Atoi(input.Text)

				switch valueInput {
				//Ace, Two, Three,..., Ten, Jack, Queen, King
				case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12:
					valid = false

				default:
					fmt.Println("Invalid input, please use a number from 0 to 12.")
					valid = true
				}
			}

			//Input handler for card suit
			valid = true
			for valid {
				fmt.Println("Enter a card suit (e.g. Heart, Diamond, Club, or Spade): ")
				fmt.Scan(&suitInput)
				suitInput = strings.ToLower(suitInput)

				fmt.Println(suitInput)

				/*
					scanner := bufio.NewScanner(os.Stdin)
					for {
						scanner.Scan()

					}
					suitInput = strings.ToLower(suitInput)
				*/

				//suitInput = strconv.(n)
				switch suitInput {
				case "heart", "diamond", "club", "spade", "hearts", "diamonds", "clubs", "spades":
					valid = false

				default:
					fmt.Println("Invalid input, please use one of the four provided suits.")
					valid = true
				}
			}

			fmt.Println(deck.GetCardIndex(currDeck, valueInput, suitInput))

		case "4":
			//Get card name from hand
			fmt.Println("Choose a card from the hand (#): ")
			valid := true

			//Reads out each card in the hand with its selection number, value, and suit
			for i := 0; i < len(currentHand.ActualHand); i++ {
				fmt.Println("Card #" + strconv.Itoa(i+1) + ": val - " + strconv.Itoa(currentHand.ActualHand[i].Val) + "  suit - " + currentHand.ActualHand[i].Suit)
			}

			for valid {
				var selection int
				valid = false
				fmt.Scanln(selection)

				switch selection {
				case 1:
					card.GetCardName(currentHand.ActualHand[0])
				case 2:
					card.GetCardName(currentHand.ActualHand[1])
				case 3:
					card.GetCardName(currentHand.ActualHand[2])
				case 4:
					card.GetCardName(currentHand.ActualHand[3])
				case 5:
					card.GetCardName(currentHand.ActualHand[4])
				default:
					fmt.Println("Invalid input, please select the number of the card from one of the " + strconv.Itoa(len(currentHand.ActualHand)) + " cards in the hand.")
					valid = true
				}
			}

		case "5":
			fmt.Println("Welcome to Helping Hand's backend environment")
			fmt.Println("To start the program, press enter")
			fmt.Println("Options:\n1: Generate deck\n2: Add card to hand\n3: Find card index in deck\n4: Get card name from hand\n5: Clear screen")

		default:
			menuOn = true
		}
	}
	r.Run("0.0.0.0:5000")
}
