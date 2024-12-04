package item

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: MEDIUM: Add Endpoint to get all borrow list by id

type BorrowListController struct {
}

func NewBorrowListController(r *gin.RouterGroup) *BorrowListController {
	group := r.Group("/borrow-list")

	ctl := BorrowListController{}
	group.GET("/ping", ctl.Ping)

	return &ctl
}

func (b *BorrowListController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"foo": "bar",
	})
}
