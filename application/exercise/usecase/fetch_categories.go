package usecase

import (
	"context"
	"healthRoutine/application/domain/exercise"
)

var _ exercise.FetchCategoriesUseCase = (*fetchCategoriesUseCaseImpl)(nil)

func FetchCategoriesUseCase(repo exercise.Repository) exercise.FetchCategoriesUseCase {
	return &fetchCategoriesUseCaseImpl{
		Repository: repo,
	}
}

type fetchCategoriesUseCaseImpl struct {
	exercise.Repository
}

func (u *fetchCategoriesUseCaseImpl) Use(ctx context.Context) ([]exercise.ExerciseCategoryModel, error) {
	return u.Repository.FetchCategories(ctx)
}
