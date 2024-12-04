package borrowlist

import (
	"net/http"

	"github.com/Lunarisnia/inventory-manager/database/repo"
	"github.com/gin-gonic/gin"
)

// TODO: MEDIUM: Add Endpoint to get all borrow list by id

type BorrowListController struct {
	repository *repo.Queries
}

func NewBorrowListController(r *gin.RouterGroup, repository *repo.Queries) *BorrowListController {
	group := r.Group("/borrow-list")

	ctl := BorrowListController{
		repository: repository,
	}
	group.GET("/ping", ctl.Ping)

	return &ctl
}

func (b *BorrowListController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"foo": "bar",
	})
}
