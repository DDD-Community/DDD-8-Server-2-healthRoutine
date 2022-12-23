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

func (u *deleteHealthUseCase) Use(ctx context.Context, userId uuid.UUID, id int64, time int64) (err error) {
	err = u.repo.DeleteHealth(ctx, userId, id, time)
	if err != nil {
		return
	}

	return
}
