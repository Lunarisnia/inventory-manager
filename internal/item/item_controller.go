package item

import (
	"net/http"

	"github.com/Lunarisnia/inventory-manager/database/repo"
	"github.com/gin-gonic/gin"
)

// TODO: EASY: Add List item endpoint

type ItemController struct {
	repository *repo.Queries
}

func NewItemController(r *gin.RouterGroup, repository *repo.Queries) *ItemController {
	group := r.Group("/item")

	ctl := ItemController{
		repository: repository,
	}
	group.GET("/ping", ctl.Ping)
	group.POST("/borrow", ctl.Borrow)
	group.GET("/bastard", ctl.Bastard)

	return &ctl
}

func (i *ItemController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"foo": "bar",
	})
}

func (i *ItemController) Borrow(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Ok",
	})
}
func (i *ItemController) Bastard(c *gin.Context) {
	c.JSON(200, gin.H{
		"hey_you_fuckYOU": "yes_i do",
	})
}
