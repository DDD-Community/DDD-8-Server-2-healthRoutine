package user

import (
	"context"
	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, nickname, email, password string) error
	GetById(ctx context.Context, id uuid.UUID) (*DomainModel, error)
	GetByEmail(ctx context.Context, email string) (*DomainModel, error)
	CheckExistsByEmail(ctx context.Context, email string) (bool, error)
	CheckExistsByNickname(ctx context.Context, nickname string) (bool, error)
	UpdateProfileById(ctx context.Context, id uuid.UUID, nickname, url string) error
}
