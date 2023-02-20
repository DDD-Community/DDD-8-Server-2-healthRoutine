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

func GetTodayWaterByUserIdUseCase(repo exercise.Repository) exercise.GetTodayWaterByUserIdUseCase {
	return &getTodayWaterByUserIdUseCaseImpl{
		Repository: repo,
	}
}

type getTodayWaterByUserIdUseCaseImpl struct {
	exercise.Repository
}

type getWaterByUserIdUseCaseImpl struct {
	exercise.Repository
}

func (u *getWaterByUserIdUseCaseImpl) Use(ctx context.Context, userId uuid.UUID, y, m, d int) (entity.Water, error) {
	return u.Repository.GetWaterByUserId(ctx, userId, y, m, d)
}

func (u *getTodayWaterByUserIdUseCaseImpl) Use(ctx context.Context, userId uuid.UUID) (entity.Water, error) {
	return u.Repository.GetTodayWaterByUserId(ctx, userId)
}
