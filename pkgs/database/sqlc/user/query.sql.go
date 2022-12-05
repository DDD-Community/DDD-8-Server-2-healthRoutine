// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: query.sql

package user

import (
	"context"

	"github.com/google/uuid"
)

const checkExistsByEmail = `-- name: CheckExistsByEmail :one
SELECT EXISTS(
    SELECT id, nickname, email, password, profile_img, created_at, updated_at FROM users
    WHERE email = ?
           ) AS isExist
`

func (q *Queries) CheckExistsByEmail(ctx context.Context, email string) (bool, error) {
	row := q.queryRow(ctx, q.checkExistsByEmailStmt, checkExistsByEmail, email)
	var isexist bool
	err := row.Scan(&isexist)
	return isexist, err
}

const checkExistsByNickname = `-- name: CheckExistsByNickname :one
SELECT EXISTS(
   SELECT id, nickname, email, password, profile_img, created_at, updated_at FROM users
   WHERE nickname = ?
           ) AS isExsist
`

func (q *Queries) CheckExistsByNickname(ctx context.Context, nickname string) (bool, error) {
	row := q.queryRow(ctx, q.checkExistsByNicknameStmt, checkExistsByNickname, nickname)
	var isexsist bool
	err := row.Scan(&isexsist)
	return isexsist, err
}

const create = `-- name: Create :exec
INSERT INTO users (id, nickname, email, password, profile_img, created_at, updated_at) VALUES (?,?,?,?,?,?,?)
`

type CreateParams struct {
	ID         uuid.UUID
	Nickname   string
	Email      string
	Password   string
	ProfileImg string
	CreatedAt  int64
	UpdatedAt  int64
}

func (q *Queries) Create(ctx context.Context, arg CreateParams) error {
	_, err := q.exec(ctx, q.createStmt, create,
		arg.ID,
		arg.Nickname,
		arg.Email,
		arg.Password,
		arg.ProfileImg,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const getByEmail = `-- name: GetByEmail :one
SELECT id, nickname, email, password, profile_img, created_at, updated_at from users
WHERE email = ?
`

func (q *Queries) GetByEmail(ctx context.Context, email string) (User, error) {
	row := q.queryRow(ctx, q.getByEmailStmt, getByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Nickname,
		&i.Email,
		&i.Password,
		&i.ProfileImg,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
