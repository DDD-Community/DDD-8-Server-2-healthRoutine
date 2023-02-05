package usecase

import (
	"context"
	"github.com/google/uuid"
	"healthRoutine/application/domain/exercise"
)

var _ exercise.FetchExerciseByCategoryIdUseCase = (*fetchExerciseByCategoryIdUseCaseImpl)(nil)

func FetchExerciseByCategoryIdUseCase(repo exercise.Repository) exercise.FetchExerciseByCategoryIdUseCase {
	return &fetchExerciseByCategoryIdUseCaseImpl{
		Repository: repo,
	}
}

type fetchExerciseByCategoryIdUseCaseImpl struct {
	exercise.Repository
}

func (u *fetchExerciseByCategoryIdUseCaseImpl) Use(ctx context.Context, userId uuid.UUID, categoryId int64) ([]exercise.ExerciseModel, error) {
	return u.Repository.FetchExerciseByCategoryId(ctx, userId, categoryId)
}
