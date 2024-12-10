package borrowlistmodels

type BorrowNewItem struct {
	ItemID int32 `json:"item_id"`
}

type ReturnItemRequest struct {
	ItemID int32 `json:"item_id"`
}
