package main

import (
	"example/web-service-gin/deck"
	"example/web-service-gin/hand"
	"example/web-service-gin/handler"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	currDeck := deck.NewDeck()
	fmt.Println(deck.GetCardIndex(currDeck, 0, "Heart"))
	currentHand := hand.NewHand("initial")
	hand.AddCardHandRandom(currentHand)
	fmt.Println("current length of hand is: ")
	fmt.Println(len(currentHand.ActualHand))
	//user
	currUserHand := hand.New()
	currUserProb := deck.New()

	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/ping", handler.PingGet())
		api.GET("/hand", handler.HandGet(currUserHand))
		api.GET("/prob", handler.ProbGet(currUserProb))
		api.POST("/hand", handler.HandPost(currUserHand, currUserHand, currDeck, currUserProb)) //deck object
	}
	r.Run("0.0.0.0:5000")
}
