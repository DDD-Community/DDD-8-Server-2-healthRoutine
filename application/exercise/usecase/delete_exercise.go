package usecase

import (
	"context"
	"github.com/google/uuid"
	"healthRoutine/application/domain/exercise"
)

var _ exercise.DeleteExerciseUseCase = (*deleteExerciseUseCase)(nil)

func DeleteExerciseUseCase(repo exercise.Repository) exercise.DeleteExerciseUseCase {
	return &deleteExerciseUseCase{
		repo: repo,
	}
}

type deleteExerciseUseCase struct {
	repo exercise.Repository
}

func (u *deleteExerciseUseCase) Use(ctx context.Context, id int64, userId uuid.UUID) (err error) {
	resp, err := u.repo.GetExerciseById(ctx, id)
	if err != nil {
		return
	}

	if userId == *resp.UserID {
		err = u.repo.DeleteExercise(ctx, id, userId)
		if err != nil {
			return
		}
	} else {
		err = exercise.ErrNotMatchUserId
		return
	}

	return nil
}
