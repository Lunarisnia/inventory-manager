-- name: ListActiveBorrowListByItemID :many
SELECT * FROM borrow_lists WHERE item_id = $1 AND returned_at IS NULL ORDER BY borrow_at DESC;

-- name: ListActiveBorrowListByUserID :many
SELECT sqlc.embed(b), sqlc.embed(i)
FROM borrow_lists as b INNER JOIN items as i ON b.item_id = i.id
WHERE user_id = $1 AND returned_at IS NULL ORDER BY borrow_at DESC;

-- name: ListAllBorrowListByUserID :many
SELECT * FROM borrow_lists WHERE user_id = $1 ORDER BY borrow_at DESC;

-- name: UpdateBorrowListReturnedAt :exec
UPDATE borrow_lists
	set returned_at = $3
WHERE id = (SELECT (id) FROM borrow_lists WHERE borrow_lists.user_id = $1 AND borrow_lists.item_id = $2 AND borrow_lists.returned_at is null LIMIT 1)
RETURNING *;

-- name: CreateBorrowList :one
INSERT INTO borrow_lists (
	user_id, item_id, borrow_at, created_at, updated_at
) VALUES (
	$1, $2, $3, $4, $5
) RETURNING *;

-- name: DeleteBorrowList :exec
DELETE FROM borrow_lists WHERE id = $1;

