package item

import (
	"net/http"
	"time"

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
	group.GET("/itemdetail", ctl.GetItemDetail)
	group.POST("/newitem", ctl.CreateNewItem)
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

func (i *ItemController) GetItemDetail(c *gin.Context) {
	c.JSON(200, gin.H{
		"TestItem": "item1",
	})
}

func (i *ItemController) CreateNewItem(c *gin.Context) {
	i.repository.CreateItem(c.Request.Context(), repo.CreateItemParams{
		Name: "Teuku",
		Image: "TestImage",
		Quantity: 1,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	})
	i.repository.ChangePassword(c.Request.Context(), repo.ChangePasswordParams{
		ID: 1,
		Password: "usergoblok",
	})
	c.JSON(200, gin.H{
		"TestCreateItem": "item1",
	})
}
