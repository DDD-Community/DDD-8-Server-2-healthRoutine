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

func (u *fetchTodayExerciseByUserIdUseCaseImpl) Use(ctx context.Context, userId uuid.UUID, time int64) (res []exercise.FetchTodayExerciseResult, err error) {
	resp, err := u.Repository.FetchTodayExerciseByUserId(ctx, userId, time)
	if err != nil {
		return
	}

	res = make([]exercise.FetchTodayExerciseResult, 0, len(resp))
	for _, v := range resp {
		res = append(res, exercise.FetchTodayExerciseResult{
			Id:              v.ID,
			ExerciseSubject: v.ExerciseSubject,
			CategorySubject: v.CategorySubject,
			Weight:          v.Weight,
			Set:             v.Set,
			Reps:            v.Reps,
		})
	}

	return
}
