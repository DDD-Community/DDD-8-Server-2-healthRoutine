package user

import "context"

type Repository interface {
	Create(ctx context.Context, nickname, email, password string) error
	GetByEmail(ctx context.Context, email string) (*DomainModel, error)
}
