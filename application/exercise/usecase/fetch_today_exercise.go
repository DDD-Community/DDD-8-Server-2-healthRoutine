package usecase

import (
	"context"
	"github.com/google/uuid"
	"healthRoutine/application/domain/exercise"
)

func FetchTodayExerciseByUserIdUseCase(repo exercise.Repository) exercise.FetchTodayExerciseByUserIdUseCase {
	return &fetchTodayExerciseByUserIdUseCaseImpl{
		Repository: repo,
	}
}

type fetchTodayExerciseByUserIdUseCaseImpl struct {
	exercise.Repository
}

func (u *fetchTodayExerciseByUserIdUseCaseImpl) Use(ctx context.Context, userId uuid.UUID, time int64) (res []exercise.FetchTodayExerciseParams, err error) {
	resp, err := u.Repository.FetchTodayExerciseByUserId(ctx, userId, time)
	if err != nil {
		return
	}

	res = make([]exercise.FetchTodayExerciseParams, 0, len(resp))
	for _, v := range resp {
		res = append(res, exercise.FetchTodayExerciseParams{
			Subject: v.Subject,
			Count:   v.Count,
		})
	}

	return
}
