package handler

import (
	"example/web-service-gin/card"
	"example/web-service-gin/deck"
	"example/web-service-gin/hand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handPostRequest struct {
	Val   int    `json:"Val"`
	Suit  string `json:"Suit"`
	Index int    `json:"Index"`
}

func HandPost(currUserHand hand.Adder, currUserHandGet hand.Getter, currDeck deck.Deck, currUserProb deck.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {

		requestBody := handPostRequest{}
		c.Bind(&requestBody)

		item := card.Card{
			Val:   requestBody.Val,
			Suit:  requestBody.Suit,
			Index: requestBody.Index,
		}
		currUserHand.Add(item)
		//currUserHand.ActualHand = append(currUserHand.ActualHand, item)

		//Update Probibilities
		arr := currUserHandGet.GetAll()
		deck.UpdateProb(arr, currDeck, currUserProb.GetAll())

		c.Status(http.StatusNoContent)
	}
}
