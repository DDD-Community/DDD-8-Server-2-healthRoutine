package usecase

import (
	"context"
	"github.com/google/uuid"
	"healthRoutine/application/domain/exercise"
	entity "healthRoutine/pkgs/database/sqlc/exercise"
)

func GetWaterByUserIdUseCase(repo exercise.Repository) exercise.GetWaterByUserIdUseCase {
	return &getWaterByUserIdUseCaseImpl{
		Repository: repo,
	}
}

type getWaterByUserIdUseCaseImpl struct {
	exercise.Repository
}

func (u *getWaterByUserIdUseCaseImpl) Use(ctx context.Context, userId uuid.UUID) (entity.Water, error) {
	return u.Repository.GetWaterByUserId(ctx, userId)
}
