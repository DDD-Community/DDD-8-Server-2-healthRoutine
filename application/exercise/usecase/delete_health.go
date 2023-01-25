package usecase

import (
	"context"
	"github.com/google/uuid"
	"healthRoutine/application/domain/exercise"
)

var _ exercise.DeleteHealthUseCase = (*deleteHealthUseCase)(nil)

func DeleteHealthUseCase(repo exercise.Repository) exercise.DeleteHealthUseCase {
	return &deleteHealthUseCase{
		repo: repo,
	}
}

type deleteHealthUseCase struct {
	repo exercise.Repository
}

func (u *deleteHealthUseCase) Use(ctx context.Context, id uuid.UUID) (err error) {
	err = u.repo.DeleteHealth(ctx, id)
	if err != nil {
		return
	}

	return
}
