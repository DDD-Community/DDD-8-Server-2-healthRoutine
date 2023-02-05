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

type LatestBadge struct {
	Index   int64  `json:"index"`
	Subject string `json:"subject"`
}

type GetBadge struct {
	ExerciseStart     bool        `json:"exerciseStart"`
	ExerciseHappy     bool        `json:"exerciseHappy"`
	ExerciseHolic     bool        `json:"exerciseHolic"`
	ExerciseMaster    bool        `json:"exerciseMaster"`
	ExerciseChampion  bool        `json:"exerciseChampion"`
	SincerityJunior   bool        `json:"sincerityJunior"`
	SincerityPro      bool        `json:"sincerityPro"`
	SincerityMaster   bool        `json:"sincerityMaster"`
	SincerityChampion bool        `json:"sincerityChampion"`
	DrinkHoneyHoney   bool        `json:"drinkHoneyHoney"`
	DrinkBulkUpBulkUp bool        `json:"drinkBulkUpBulkUp"`
	DrinkHippo        bool        `json:"drinkHippo"`
	LatestBadge       LatestBadge `json:"latestBadge"`
}

type GetBadgeUseCase interface {
	Use(ctx context.Context, userId uuid.UUID) (*GetBadge, error)
}

type WithdrawalUseCase interface {
	Use(ctx context.Context, userId uuid.UUID) error
}
