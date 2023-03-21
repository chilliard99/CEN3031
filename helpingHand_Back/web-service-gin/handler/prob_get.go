package handler

import (
	"example/web-service-gin/deck"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProbGet(currUserProb deck.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := currUserProb.GetAll()
		//fmt.Println("hello im here " + results[0].Suit)
		c.JSON(http.StatusOK, results)
	}
}
