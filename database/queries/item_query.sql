-- name: GetItem :one
SELECT * FROM items WHERE id = $1 LIMIT 1;

-- name: CreateItem :one
INSERT INTO items (name, image, quantity, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: UpdateItemQuantity :exec
UPDATE items 
	set quantity = $2
WHERE id = $1
RETURNING *;

-- name: DeleteItem :exec
DELETE FROM items
WHERE id = $1;

-- name: ListItem :many
SELECT * FROM items;
