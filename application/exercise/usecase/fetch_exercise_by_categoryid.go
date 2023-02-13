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

// Use
// TODO: refactor
func (u *fetchExerciseByCategoryIdUseCaseImpl) Use(ctx context.Context, userId uuid.UUID) (res []exercise.FetchExerciseResult, err error) {
	resp, err := u.Repository.FetchCategories(ctx)
	if err != nil {
		return
	}

	res = make([]exercise.FetchExerciseResult, 0, len(resp))
	for _, v := range resp {
		exResp, ferr := u.Repository.FetchExerciseByCategoryId(ctx, userId, v.ExerciseCategory.ID)
		if ferr != nil {
			return nil, ferr
		}
		res = append(res, exercise.FetchExerciseResult{
			Id:       v.ExerciseCategory.ID,
			Subject:  v.ExerciseCategory.Subject,
			Exercise: exResp,
		})
	}

	return
}
