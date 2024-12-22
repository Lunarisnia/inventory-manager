package borrowlistmodels

import "github.com/Lunarisnia/inventory-manager/database/repo"

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
	BorrowAt   int64  `json:"borrow_at"`
	ReturnedAt *int64 `json:"returned_at"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
}

func (b *BorrowListResponse) ConvertRaw(r repo.BorrowList) {
	b.ReturnedAt = nil
	b.ID = r.ID
	b.UserID = r.UserID
	b.ItemID = r.ItemID
	b.BorrowAt = r.BorrowAt
	if r.ReturnedAt.Valid {
		b.ReturnedAt = &r.ReturnedAt.Int64
	}
	b.CreatedAt = r.CreatedAt
	b.UpdatedAt = r.UpdatedAt
}
