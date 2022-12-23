package usecase

import (
	"context"
	"healthRoutine/application/domain/exercise"
)

var _ exercise.CreateExerciseUseCase = (*createExerciseUseCaseImpl)(nil)

func CreateExerciseUseCase(repo exercise.Repository) exercise.CreateExerciseUseCase {
	return &createExerciseUseCaseImpl{
		Repository: repo,
	}
}

type createExerciseUseCaseImpl struct {
	exercise.Repository
}

func (u *createExerciseUseCaseImpl) Use(ctx context.Context, params exercise.CreateExerciseParams) error {
	return u.Repository.Create(ctx, params.UserId, params.ExerciseId, params.Weight, params.Set, params.Minute)
}
