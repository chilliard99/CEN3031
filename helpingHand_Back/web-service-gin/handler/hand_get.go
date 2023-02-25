package handler

import (
	"example/web-service-gin/hand"
	//"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandGet(currUserHand hand.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := currUserHand.GetAll()
		//fmt.Println("hello im here " + results[0].Suit)
		c.JSON(http.StatusOK, results)
	}
}
