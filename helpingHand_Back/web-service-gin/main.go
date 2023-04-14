package main

import (
	"example/web-service-gin/deck"
	"example/web-service-gin/hand"
	"example/web-service-gin/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	CurrDeck := deck.NewDeck()

	//Lines from the very beginning of development. I believe they're all non-functional.
	//fmt.Println(deck.GetCardIndex(currDeck, 0, "Heart"))
	//currentHand := hand.NewHand("initial")
	//hand.AddCardHandRandom(currentHand)
	//fmt.Println("current length of hand is: ")
	//fmt.Println(len(currentHand.ActualHand))

	//user
	CurrUserHand := hand.New()
	CurrUserProb := deck.New()

	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/ping", handler.PingGet())
		api.GET("/hand", handler.HandGet(CurrUserHand))
		api.GET("/prob", handler.ProbGet(CurrUserProb))
		api.POST("/hand", handler.HandPost(CurrUserHand, CurrUserHand, CurrDeck, CurrUserProb)) //deck object
		api.GET("/removeAll", handler.HandDelete(CurrUserHand, CurrUserProb))
	}
	r.Run("0.0.0.0:5000")
}
