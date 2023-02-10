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

type UpdateProfileParams struct {
	Id       uuid.UUID
	Nickname string
}

type UpdateProfileUseCase interface {
	Use(ctx context.Context, params UpdateProfileParams) error
}

type UploadTemporaryProfileParams struct {
	Id            uuid.UUID
	Filename      string
	ContentType   string
	ContentLength int64
	ProfileImg    io.Reader
}

type UploadTemporaryProfileUseCase interface {
	Use(ctx context.Context, params UploadTemporaryProfileParams) (string, error)
}

type GetBadgeResult struct {
	MyBadge      []string `json:"myBadge"`
	WaitingBadge []string `json:"waitingBadge"`
	LatestBadge  string   `json:"latestBadge"`
}

type GetBadgeUseCase interface {
	Use(ctx context.Context, userId uuid.UUID) (*GetBadgeResult, error)
}

type WithdrawalUseCase interface {
	Use(ctx context.Context, userId uuid.UUID) error
}
