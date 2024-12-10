// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: item_query.sql

package repo

import (
	"context"
)

const createItem = `-- name: CreateItem :one
INSERT INTO items (name, image, quantity, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5) RETURNING id, name, image, quantity, created_at, updated_at
`

type CreateItemParams struct {
	Name      string
	Image     string
	Quantity  int32
	CreatedAt int64
	UpdatedAt int64
}

func (q *Queries) CreateItem(ctx context.Context, arg CreateItemParams) (Item, error) {
	row := q.db.QueryRow(ctx, createItem,
		arg.Name,
		arg.Image,
		arg.Quantity,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Image,
		&i.Quantity,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteItem = `-- name: DeleteItem :exec
DELETE FROM items
WHERE id = $1
`

func (q *Queries) DeleteItem(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteItem, id)
	return err
}

const getItem = `-- name: GetItem :one
SELECT id, name, image, quantity, created_at, updated_at FROM items WHERE id = $1 LIMIT 1
`

func (q *Queries) GetItem(ctx context.Context, id int32) (Item, error) {
	row := q.db.QueryRow(ctx, getItem, id)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Image,
		&i.Quantity,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listItem = `-- name: ListItem :many
SELECT id, name, image, quantity, created_at, updated_at FROM items
`

func (q *Queries) ListItem(ctx context.Context) ([]Item, error) {
	rows, err := q.db.Query(ctx, listItem)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Image,
			&i.Quantity,
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

const updateItemQuantity = `-- name: UpdateItemQuantity :exec
UPDATE items 
	set quantity = $2
WHERE id = $1
RETURNING id, name, image, quantity, created_at, updated_at
`

type UpdateItemQuantityParams struct {
	ID       int32
	Quantity int32
}

func (q *Queries) UpdateItemQuantity(ctx context.Context, arg UpdateItemQuantityParams) error {
	_, err := q.db.Exec(ctx, updateItemQuantity, arg.ID, arg.Quantity)
	return err
}
