package handler

import (
	//"example/web-service-gin/card"
	"example/web-service-gin/deck"
	"example/web-service-gin/hand"

	//"example/web-service-gin/main.go"
	//"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandDelete(currUserHand hand.Getter, currUserProb deck.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		//fmt.Println(len(currDeck))
		currUserHand.Reset()
		result := currUserHand.GetAll()
		//result := *currUserHand.GetAll()
		//fmt.Println(len(currDeck))
		c.JSON(http.StatusOK, result)
		//c.Status(http.StatusNoContent)
	}
}
