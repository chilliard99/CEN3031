package handler

import (
	c "example/web-service-gin/card"
	"example/web-service-gin/hand"
	"sort"

	//"example/web-service-gin/deck"

	"net/http"

	"github.com/gin-gonic/gin"
)

type ByIndex []c.Card

func (a ByIndex) Len() int           { return len(a) }
func (a ByIndex) Less(i, j int) bool { return a[i].Index < a[j].Index }
func (a ByIndex) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func HandGet(currUserHand hand.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := currUserHand.GetAll()
		sort.Sort(ByIndex(results))
		c.JSON(http.StatusOK, results)
	}
}
