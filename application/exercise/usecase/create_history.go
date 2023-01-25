package usecase

import (
	"context"
	"healthRoutine/application/domain/exercise"
)

var _ exercise.CreateHistoryUseCase = (*createHistoryUseCaseImpl)(nil)

func CreateHistoryUseCase(repo exercise.Repository) exercise.CreateHistoryUseCase {
	return &createHistoryUseCaseImpl{
		Repository: repo,
	}
}

type createHistoryUseCaseImpl struct {
	exercise.Repository
}

func (u *createHistoryUseCaseImpl) Use(ctx context.Context, params exercise.CreateHistoryParams) error {
	return u.Repository.Create(ctx, params.UserId, params.ExerciseId, params.Weight, params.Reps, params.Set)
}
