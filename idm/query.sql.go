// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package idm

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO idm_users (username, password, first_name, last_name, created_at, last_modified_at, deleted_at)
VALUES ($1, $2, $3, $4, NOW() AT TIME ZONE 'utc', NOW() AT TIME ZONE 'utc', NULL)
RETURNING uuid
`

type CreateUserParams struct {
	Username  string
	Password  []byte
	FirstName string
	LastName  sql.NullString
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.Exec(ctx, createUser,
		arg.Username,
		arg.Password,
		arg.FirstName,
		arg.LastName,
	)
	return err
}

const deleteUserByUsername = `-- name: DeleteUserByUsername :exec
DELETE FROM idm_users
WHERE username = $1
`

func (q *Queries) DeleteUserByUsername(ctx context.Context, username string) error {
	_, err := q.db.Exec(ctx, deleteUserByUsername, username)
	return err
}

const findUserByUsername = `-- name: FindUserByUsername :one
SELECT uuid, username, password
FROM idm_users
WHERE username = $1
`

type FindUserByUsernameRow struct {
	Uuid     uuid.UUID
	Username string
	Password []byte
}

func (q *Queries) FindUserByUsername(ctx context.Context, username string) (FindUserByUsernameRow, error) {
	row := q.db.QueryRow(ctx, findUserByUsername, username)
	var i FindUserByUsernameRow
	err := row.Scan(&i.Uuid, &i.Username, &i.Password)
	return i, err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE idm_users
SET username = $2, password = $3, first_name = $4, last_name = $5, last_modified_at = NOW() AT TIME ZONE 'utc'
WHERE uuid = $1
`

type UpdateUserParams struct {
	Uuid      uuid.UUID
	Username  string
	Password  []byte
	FirstName string
	LastName  sql.NullString
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.Exec(ctx, updateUser,
		arg.Uuid,
		arg.Username,
		arg.Password,
		arg.FirstName,
		arg.LastName,
	)
	return err
}
