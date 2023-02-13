package handler

import (
	"example/web-service-gin/hand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandGet(currUserHand hand.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := currUserHand.GetAll()
		c.JSON(http.StatusOK, results)
	}
}
