package borrowlist

import (
	"net/http"
	"time"

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
	group.GET("/getallborrowlist", ctl.GetAllBorrowList)
	group.POST("/createnewborrowlist", ctl.CreateNewBorrowList)

	return &ctl
}

func (b *BorrowListController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"foo": "bar",
	})
}

func (b *BorrowListController) GetAllBorrowList(c *gin.Context) {
	resp, _ := b.repository.ListAllBorrowListByUserID(c.Request.Context(), 2)
	c.JSON(http.StatusOK, gin.H{
		"foo": resp,
	})
}
func (b *BorrowListController) CreateNewBorrowList(c *gin.Context) {
	b.repository.CreateBorrowList(c.Request.Context(), repo.CreateBorrowListParams{
		UserID:    1,
		ItemID:    1,
		BorrowAt:  1,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	})
	c.JSON(200, gin.H{
		"foo": "bar",
	})
}
