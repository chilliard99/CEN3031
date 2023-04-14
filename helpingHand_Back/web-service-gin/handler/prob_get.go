package handler

import (
	"example/web-service-gin/deck"
	"example/web-service-gin/hand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProbGet(currUserHand hand.Getter, currDeck deck.Deck, currUserProb deck.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := currUserProb.GetAll()
		arr := currUserHand.GetAll()
		deck.UpdateProb(arr, currDeck, currUserProb.GetAll())
		//fmt.Println("hello im here " + results[0].Suit)
		c.JSON(http.StatusOK, results)
	}
}
