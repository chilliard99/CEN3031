package handler

import (
	"example/web-service-gin/card"
	"example/web-service-gin/deck"
	"example/web-service-gin/hand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handPostRequest struct {
	Val  int    `json:"Val"`
	Suit string `json:"Suit"`
}

func HandPost(currUserHand hand.Adder, currUserHandGet hand.Getter, currDeck deck.Deck) gin.HandlerFunc {
	return func(c *gin.Context) {

		requestBody := handPostRequest{}
		c.Bind(&requestBody)

		item := card.Card{
			Val:  requestBody.Val,
			Suit: requestBody.Suit,
		}
		currUserHand.Add(item)
		//currUserHand.ActualHand = append(currUserHand.ActualHand, item)

		//Update Probibilities
		arr := currUserHandGet.GetAll()
		deck.UpdateProb(arr, currDeck)

		c.Status(http.StatusNoContent)
	}
}
