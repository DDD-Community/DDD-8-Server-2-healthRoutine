package user

import (
	"context"
	"github.com/google/uuid"
	entity "healthRoutine/pkgs/database/sqlc/user"
)

type Repository interface {
	Create(ctx context.Context, nickname, email, password string) error
	GetById(ctx context.Context, id uuid.UUID) (*DomainModel, error)
	GetByEmail(ctx context.Context, email string) (*DomainModel, error)
	GetNicknameById(ctx context.Context, userId uuid.UUID) (string, error)
	CheckExistsByEmail(ctx context.Context, email string) (bool, error)
	CheckExistsByNickname(ctx context.Context, nickname string) (bool, error)
	UpdateProfileById(ctx context.Context, id uuid.UUID, nickname, url string) error
	CreateBadge(ctx context.Context, userId uuid.UUID, badgeId []int64) error
	GetBadgeByUserId(ctx context.Context) ([]entity.Badge, error)
	GetLatestBadgeByUserId(ctx context.Context, userId uuid.UUID) (entity.GetLatestBadgeByUserIdRow, error)
	ExistsBadgeByUserId(ctx context.Context, userId uuid.UUID, badgeId int64) (bool, error)
	DeleteUserById(ctx context.Context, userId uuid.UUID) error
}
