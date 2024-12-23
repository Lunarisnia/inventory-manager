// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: borrow_list_query.sql

package repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createBorrowList = `-- name: CreateBorrowList :one
INSERT INTO borrow_lists (
	user_id, item_id, borrow_at, created_at, updated_at
) VALUES (
	$1, $2, $3, $4, $5
) RETURNING id, user_id, item_id, borrow_at, returned_at, created_at, updated_at
`

type CreateBorrowListParams struct {
	UserID    int32
	ItemID    int32
	BorrowAt  int64
	CreatedAt int64
	UpdatedAt int64
}

func (q *Queries) CreateBorrowList(ctx context.Context, arg CreateBorrowListParams) (BorrowList, error) {
	row := q.db.QueryRow(ctx, createBorrowList,
		arg.UserID,
		arg.ItemID,
		arg.BorrowAt,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i BorrowList
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ItemID,
		&i.BorrowAt,
		&i.ReturnedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteBorrowList = `-- name: DeleteBorrowList :exec
DELETE FROM borrow_lists WHERE id = $1
`

func (q *Queries) DeleteBorrowList(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteBorrowList, id)
	return err
}

const listActiveBorrowListByItemID = `-- name: ListActiveBorrowListByItemID :many
SELECT id, user_id, item_id, borrow_at, returned_at, created_at, updated_at FROM borrow_lists WHERE item_id = $1 AND returned_at IS NULL ORDER BY borrow_at DESC
`

func (q *Queries) ListActiveBorrowListByItemID(ctx context.Context, itemID int32) ([]BorrowList, error) {
	rows, err := q.db.Query(ctx, listActiveBorrowListByItemID, itemID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []BorrowList
	for rows.Next() {
		var i BorrowList
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ItemID,
			&i.BorrowAt,
			&i.ReturnedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listActiveBorrowListByUserID = `-- name: ListActiveBorrowListByUserID :many
SELECT b.id, b.user_id, b.item_id, b.borrow_at, b.returned_at, b.created_at, b.updated_at, i.id, i.name, i.image, i.quantity, i.created_at, i.updated_at
FROM borrow_lists as b INNER JOIN items as i ON b.item_id = i.id
WHERE user_id = $1 AND returned_at IS NULL ORDER BY borrow_at DESC
`

type ListActiveBorrowListByUserIDRow struct {
	BorrowList BorrowList
	Item       Item
}

func (q *Queries) ListActiveBorrowListByUserID(ctx context.Context, userID int32) ([]ListActiveBorrowListByUserIDRow, error) {
	rows, err := q.db.Query(ctx, listActiveBorrowListByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListActiveBorrowListByUserIDRow
	for rows.Next() {
		var i ListActiveBorrowListByUserIDRow
		if err := rows.Scan(
			&i.BorrowList.ID,
			&i.BorrowList.UserID,
			&i.BorrowList.ItemID,
			&i.BorrowList.BorrowAt,
			&i.BorrowList.ReturnedAt,
			&i.BorrowList.CreatedAt,
			&i.BorrowList.UpdatedAt,
			&i.Item.ID,
			&i.Item.Name,
			&i.Item.Image,
			&i.Item.Quantity,
			&i.Item.CreatedAt,
			&i.Item.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAllBorrowListByUserID = `-- name: ListAllBorrowListByUserID :many
SELECT id, user_id, item_id, borrow_at, returned_at, created_at, updated_at FROM borrow_lists WHERE user_id = $1 ORDER BY borrow_at DESC
`

func (q *Queries) ListAllBorrowListByUserID(ctx context.Context, userID int32) ([]BorrowList, error) {
	rows, err := q.db.Query(ctx, listAllBorrowListByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []BorrowList
	for rows.Next() {
		var i BorrowList
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ItemID,
			&i.BorrowAt,
			&i.ReturnedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateBorrowListReturnedAt = `-- name: UpdateBorrowListReturnedAt :exec
UPDATE borrow_lists
	set returned_at = $3
WHERE id = (SELECT (id) FROM borrow_lists WHERE borrow_lists.user_id = $1 AND borrow_lists.item_id = $2 AND borrow_lists.returned_at is null LIMIT 1)
RETURNING id, user_id, item_id, borrow_at, returned_at, created_at, updated_at
`

type UpdateBorrowListReturnedAtParams struct {
	UserID     int32
	ItemID     int32
	ReturnedAt pgtype.Int8
}

func (q *Queries) UpdateBorrowListReturnedAt(ctx context.Context, arg UpdateBorrowListReturnedAtParams) error {
	_, err := q.db.Exec(ctx, updateBorrowListReturnedAt, arg.UserID, arg.ItemID, arg.ReturnedAt)
	return err
}
