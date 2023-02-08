package handler

import (
	"example/web-service-gin/card"
	"example/web-service-gin/hand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handPostRequest struct {
	Val  int    `json:"val"`
	Suit string `json:"suit"`
}

func HandPost(currentHand *hand.Hand) gin.HandlerFunc {
	return func(c *gin.Context) {

		requestBody := handPostRequest{}
		c.Bind(requestBody)

		item := card.Card{
			Val:  requestBody.Val,
			Suit: requestBody.Suit,
		}

		currentHand.ActualHand = append(currentHand.ActualHand, item)

		c.Status(http.StatusNoContent)
	}
}
