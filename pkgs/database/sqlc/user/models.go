// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package user

import (
	"github.com/google/uuid"
)

type Badge struct {
	ID      int64
	Subject string
	Sub     string
}

type BadgeUser struct {
	UsersID   uuid.UUID
	BadgeID   int64
	CreatedAt int64
}

type User struct {
	ID         uuid.UUID
	Nickname   string
	Email      string
	Password   string
	ProfileImg string
	CreatedAt  int64
	UpdatedAt  int64
}
