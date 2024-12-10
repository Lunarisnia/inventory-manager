package item

import (
	"net/http"
	"time"

	"github.com/Lunarisnia/inventory-manager/database/repo"
	"github.com/Lunarisnia/inventory-manager/internal/item/itemodels"
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
	group.GET("/", ctl.ListItem)
	group.POST("/", ctl.CreateNewItem)

	return &ctl
}

func (i *ItemController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"foo": "bar",
	})
}

func (i *ItemController) ListItem(c *gin.Context) {
	items, err := i.repository.ListItem(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": items,
	})
}

func (i *ItemController) CreateNewItem(c *gin.Context) {
	newItem := itemodels.CreateNewItemRequest{}
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	item, err := i.repository.CreateItem(c.Request.Context(), repo.CreateItemParams{
		Name:      newItem.Name,
		Image:     newItem.Image,
		Quantity:  int32(newItem.Quantity),
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": item,
	})
}
