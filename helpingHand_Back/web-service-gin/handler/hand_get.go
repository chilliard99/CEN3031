package handler

import (
	"example/web-service-gin/hand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandGet(currentHand *hand.Hand) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := currentHand.ActualHand
		c.JSON(http.StatusOK, results)
	}
}
