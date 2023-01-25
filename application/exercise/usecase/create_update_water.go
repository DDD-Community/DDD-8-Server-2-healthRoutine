package usecase

import (
	"context"
	"github.com/google/uuid"
	"healthRoutine/application/domain/exercise"
)

func CreateOrUpdateWaterUseCase(repo exercise.Repository) exercise.CreateOrUpdateWaterUseCase {
	return &createOrUpdateWaterUseCaseImpl{
		Repository: repo,
	}
}

type createOrUpdateWaterUseCaseImpl struct {
	exercise.Repository
}

func (u *createOrUpdateWaterUseCaseImpl) Use(ctx context.Context, userId uuid.UUID, capacity int64) error {
	return u.Repository.CreateOrUpdateWater(ctx, userId, capacity)
}
