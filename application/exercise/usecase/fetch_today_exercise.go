package usecase

import (
	"context"
	"github.com/google/uuid"
	"healthRoutine/application/domain/exercise"
	"healthRoutine/pkgs/util/dbx"
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
			ExerciseSubject: v.Subject_2,
			CategorySubject: v.Subject,
			Weight:          dbx.ConvertInterfaceToInt64(v.Weight),
			Set:             dbx.ConvertInterfaceToInt64(v.Set),
			Count:           v.Count,
			CreatedAt:       v.CreatedAt,
		})
	}

	return
}
