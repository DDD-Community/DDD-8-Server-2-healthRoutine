package user

import (
	"context"
	"github.com/google/uuid"
	"io"
)

type SignUpParams struct {
	Nickname string
	Password string
	Email    string
}

type SignUpUseCase interface {
	Use(ctx context.Context, params SignUpParams) (*DomainModel, error)
}

type SignInUseCase interface {
	Use(ctx context.Context, email, password string) (*DomainModel, string, error)
}

type EmailValidationUseCase interface {
	Use(ctx context.Context, email string) error
}

type GetProfileUseCase interface {
	Use(ctx context.Context, id uuid.UUID) (*DomainModel, error)
}

type UpdateProfileImgParams struct {
	Id         uuid.UUID
	Filename   string
	ProfileImg io.Reader
}

type UpdateProfileImgUseCase interface {
	Use(ctx context.Context, params UpdateProfileImgParams) error
}

type UpdateNicknameUseCase interface {
	Use(ctx context.Context, id uuid.UUID, nickname string) error
}
