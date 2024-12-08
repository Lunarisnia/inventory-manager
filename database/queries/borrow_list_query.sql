-- name: ListActiveBorrowListByItemID :many
SELECT * FROM borrow_lists WHERE item_id = $1 AND returned_at IS NULL ORDER BY borrow_at DESC;

-- name: ListActiveBorrowListByUserID :many
SELECT * FROM borrow_lists WHERE user_id = $1 AND returned_at IS NULL ORDER BY borrow_at DESC;

-- name: ListAllBorrowListByUserID :many
SELECT * FROM borrow_lists WHERE user_id = $1 ORDER BY borrow_at DESC;

-- name: UpdateBorrowListReturnedAt :exec
UPDATE borrow_lists
	set returned_at = $2
WHERE id = $1
RETURNING *;

-- name: CreateBorrowList :one
INSERT INTO borrow_lists (
	user_id, item_id, borrow_at, created_at, updated_at
) VALUES (
	$1, $2, $3, $4, $5
) RETURNING *;

-- name: DeleteBorrowList :exec
DELETE FROM borrow_lists WHERE id = $1;

