// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user_query.sql

package repo

import (
	"context"
)

const changePassword = `-- name: ChangePassword :exec
UPDATE users
	SET password = $2
WHERE id = $1 RETURNING id, name, nis, password, created_at, updated_at
`

type ChangePasswordParams struct {
	ID       int32
	Password string
}

func (q *Queries) ChangePassword(ctx context.Context, arg ChangePasswordParams) error {
	_, err := q.db.Exec(ctx, changePassword, arg.ID, arg.Password)
	return err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (
	name, nis, password, created_at, updated_at
) VALUES ($1, $2, $3, $4, $5) RETURNING id, name, nis, password, created_at, updated_at
`

type CreateUserParams struct {
	Name      string
	Nis       string
	Password  string
	CreatedAt int64
	UpdatedAt int64
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Name,
		arg.Nis,
		arg.Password,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Nis,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, name, nis, password, created_at, updated_at FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Nis,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
