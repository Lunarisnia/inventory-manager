package borrowlist

import (
	"net/http"
	"time"

	"github.com/Lunarisnia/inventory-manager/database/repo"
	"github.com/Lunarisnia/inventory-manager/internal/auth"
	"github.com/Lunarisnia/inventory-manager/internal/borrowlist/borrowlistmodels"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type BorrowListController struct {
	repository *repo.Queries
}

func NewBorrowListController(r *gin.RouterGroup, repository *repo.Queries) *BorrowListController {
	group := r.Group("/borrow-list")

	ctl := BorrowListController{
		repository: repository,
	}
	group.GET("", auth.Authorized(), ctl.ListBorrowList)
	group.POST("/", auth.Authorized(), ctl.BorrowItem)
	group.POST("/return", auth.Authorized(), ctl.ReturnItem)
	group.GET("/ping", ctl.Ping)

	return &ctl
}

func (b *BorrowListController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"foo": "bar",
	})
}

func (b *BorrowListController) BorrowItem(c *gin.Context) {
	claim, ok := c.Request.Context().Value("jwt").(*auth.JWTClaim)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	borrowItem := borrowlistmodels.BorrowNewItem{}
	if err := c.BindJSON(&borrowItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request Dumbass",
		})
		return
	}

	borrowedList, err := b.repository.ListActiveBorrowListByItemID(c.Request.Context(), borrowItem.ItemID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	item, err := b.repository.GetItem(c.Request.Context(), borrowItem.ItemID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	if len(borrowedList) >= int(item.Quantity) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "item is out of stock",
		})
		return
	}

	newBorrowList, err := b.repository.CreateBorrowList(c.Request.Context(), repo.CreateBorrowListParams{
		UserID:    claim.UserID,
		ItemID:    borrowItem.ItemID,
		BorrowAt:  time.Now().Unix(),
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
		"data": newBorrowList,
	})
}

func (b *BorrowListController) ListBorrowList(c *gin.Context) {
	claim, ok := c.Request.Context().Value("jwt").(*auth.JWTClaim)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	borrowList, err := b.repository.ListActiveBorrowListByUserID(c.Request.Context(), claim.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": borrowList,
	})
}

func (b *BorrowListController) ReturnItem(c *gin.Context) {
	claim, ok := c.Request.Context().Value("jwt").(*auth.JWTClaim)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}
	returnedItem := borrowlistmodels.ReturnItemRequest{}
	if err := c.ShouldBindJSON(&returnedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	err := b.repository.UpdateBorrowListReturnedAt(c.Request.Context(), repo.UpdateBorrowListReturnedAtParams{
		UserID: claim.UserID,
		ItemID: returnedItem.ItemID,
		ReturnedAt: pgtype.Int8{
			Int64: time.Now().Unix(),
			Valid: true,
		},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "Ok",
	})
}
