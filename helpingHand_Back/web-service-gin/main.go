package main

import (
	"example/web-service-gin/deck"
	"example/web-service-gin/hand"
	"example/web-service-gin/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	CurrDeck := deck.NewDeck()

	//user
	CurrUserHand := hand.New()
	CurrUserProb := deck.New()

	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/ping", handler.PingGet())
		api.GET("/hand", handler.HandGet(CurrUserHand))
		api.GET("/prob", handler.ProbGet(CurrUserHand, CurrDeck, CurrUserProb))
		api.POST("/hand", handler.HandPost(CurrUserHand, CurrUserHand, CurrDeck, CurrUserProb)) //deck object
		api.GET("/removeAll", handler.HandDelete(CurrUserHand, CurrUserProb))
	}
	r.Run("0.0.0.0:5000")
}
