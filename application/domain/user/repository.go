package user

import "context"

type Repository interface {
	Create(ctx context.Context, nickname, email, password string) error
	GetByEmail(ctx context.Context, email string) (*DomainModel, error)
	CheckExistsByEmail(ctx context.Context, email string) (bool, error)
	CheckExistsByNickname(ctx context.Context, nickname string) (bool, error)
}
