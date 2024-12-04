package item

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: EASY: Add List item endpoint

type ItemController struct {
}

func NewItemController(r *gin.RouterGroup) *ItemController {
	group := r.Group("/item")

	ctl := ItemController{}
	group.GET("/ping", ctl.Ping)

	return &ctl
}

func (i *ItemController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"foo": "bar",
	})
}
