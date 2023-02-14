package handler

import (
	//"example/web-service-gin/card"
	"example/web-service-gin/hand"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handPostRequest struct {
	Val  int    `json:"val"`
	Suit string `json:"suit"`
}

func HandPost(currUserHand hand.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {

		requestBody := handPostRequest{}
		c.Bind(&requestBody)

		item := hand.Card{
			Val:  requestBody.Val,
			Suit: requestBody.Suit,
		}
		//delete
		fmt.Println("here")
		fmt.Println(item.Suit)

		currUserHand.Add(item)
		//currUserHand.ActualHand = append(currUserHand.ActualHand, item)

		c.Status(http.StatusNoContent)
	}
}
