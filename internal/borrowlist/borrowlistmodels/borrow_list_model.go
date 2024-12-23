package borrowlistmodels

import (
	"github.com/Lunarisnia/inventory-manager/database/repo"
)

type BorrowNewItem struct {
	ItemID int32 `json:"item_id"`
}

type ReturnItemRequest struct {
	ItemID     int32  `json:"item_id"`
	ReturnCode string `json:"return_code"`
}

type BorrowListResponse struct {
	ID         int32  `json:"id"`
	UserID     int32  `json:"user_id"`
	ItemID     int32  `json:"item_id"`
	Name       string `json:"name"`
	Image      string `json:"image"`
	BorrowAt   int64  `json:"borrow_at"`
	ReturnedAt *int64 `json:"returned_at"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
}

func (b *BorrowListResponse) ConvertRaw(r repo.ListActiveBorrowListByUserIDRow) {
	b.ReturnedAt = nil
	b.ID = r.BorrowList.ID
	b.UserID = r.BorrowList.UserID
	b.ItemID = r.BorrowList.ItemID
	b.Name = r.Item.Name
	b.BorrowAt = r.BorrowList.BorrowAt
	b.Image = r.Item.Image
	if r.BorrowList.ReturnedAt.Valid {
		b.ReturnedAt = &r.BorrowList.ReturnedAt.Int64
	}
	b.CreatedAt = r.BorrowList.CreatedAt
	b.UpdatedAt = r.BorrowList.UpdatedAt
}
