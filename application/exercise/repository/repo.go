package repository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"healthRoutine/application/domain/exercise"
	entity "healthRoutine/pkgs/database/sqlc/exercise"
	"healthRoutine/pkgs/util/timex"
	"time"
)

var _ exercise.Repository = (*repo)(nil)

func NewExerciseRepository(db *sql.DB) exercise.Repository {
	preparedQuery, err := entity.Prepare(context.Background(), db)
	if err != nil {
		panic(err)
	}
	return &repo{
		db:            db,
		query:         entity.New(db),
		preparedQuery: preparedQuery,
	}
}

type repo struct {
	db            *sql.DB
	query         *entity.Queries
	preparedQuery *entity.Queries
}

func (r *repo) Create(
	ctx context.Context,
	userId uuid.UUID,
	exerciseId int64,
	weight, set, minute int32) error {
	return r.preparedQuery.Create(ctx, entity.CreateParams{
		ID:         uuid.New(),
		UserID:     userId,
		ExerciseID: exerciseId,
		Weight:     weight,
		Set:        set,
		Minute:     minute,
		CreatedAt:  time.Now().UnixMilli(),
	})
}

func (r *repo) FetchByDateTime(ctx context.Context, userId uuid.UUID, start, end int64) (res []exercise.DomainModel, err error) {
	resp, err := r.preparedQuery.FetchByDateTime(ctx, entity.FetchByDateTimeParams{
		UserID:      userId,
		CreatedAt:   start,
		CreatedAt_2: end,
	})
	if err != nil {
		return
	}

	res = make([]exercise.DomainModel, 0, len(resp))
	for _, v := range resp {
		res = append(res, exercise.DomainModel{
			Health: entity.Health{
				ID:         v.ID,
				UserID:     v.UserID,
				ExerciseID: v.ExerciseID,
				Weight:     v.Weight,
				Set:        v.Set,
				Minute:     v.Minute,
				CreatedAt:  v.CreatedAt,
			},
		})

	}
	return
}

func (r *repo) FetchCategories(ctx context.Context) (res []exercise.ExerciseCategoryModel, err error) {
	resp, err := r.preparedQuery.FetchCategories(ctx)
	if err != nil {
		return
	}

	res = make([]exercise.ExerciseCategoryModel, 0, len(resp))
	for _, v := range resp {
		res = append(res, exercise.ExerciseCategoryModel{
			ExerciseCategory: v,
		})
	}

	return
}

func (r *repo) FetchExerciseByCategoryId(ctx context.Context, categoryId int64) (res []exercise.ExerciseModel, err error) {
	resp, err := r.preparedQuery.FetchExerciseByCategoryId(ctx, categoryId)
	if err != nil {
		return
	}

	res = make([]exercise.ExerciseModel, 0, len(resp))
	for _, v := range resp {
		res = append(res, exercise.ExerciseModel{
			Exercise: v,
		})
	}

	return
}

// TODO: bypass fix
func (r *repo) FetchTodayExerciseByUserId(ctx context.Context, userId uuid.UUID, time int64) ([]entity.FetchTodayExerciseByUserIdRow, error) {
	start, end := timex.GetDateForADayUnixMillisecond(time)

	return r.preparedQuery.FetchTodayExerciseByUserId(ctx, entity.FetchTodayExerciseByUserIdParams{
		UserID:      userId,
		CreatedAt:   start,
		CreatedAt_2: end,
	})
}
