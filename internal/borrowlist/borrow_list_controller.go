package borrowlist

import (
	"net/http"

	"github.com/Lunarisnia/inventory-manager/database/repo"
	"github.com/Lunarisnia/inventory-manager/internal/auth"
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

// WIP: THIS CRAP
func (b *BorrowListController) BorrowItem(c *gin.Context) {
	claim, ok := c.Request.Context().Value("jwt").(auth.JWTClaim)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	b.repository.CreateBorrowList(c.Request.Context(), repo.CreateBorrowListParams{
		UserID: claim.ID,
		ItemID: 2,
	})
}
