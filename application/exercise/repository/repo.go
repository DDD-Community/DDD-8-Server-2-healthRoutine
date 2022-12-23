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

// GetExerciseById
// TODO: fix bypass entity
func (r *repo) GetExerciseById(ctx context.Context, id int64) (entity.Exercise, error) {
	return r.preparedQuery.GetExerciseById(ctx, id)
}

func (r *repo) DeleteExercise(ctx context.Context, id int64, userId uuid.UUID) error {
	return r.preparedQuery.DeleteExercise(ctx, entity.DeleteExerciseParams{
		ID:     id,
		UserID: &userId,
	})
}

func (r *repo) CreateExercise(ctx context.Context, userId *uuid.UUID, categoryId int64, subject string) error {
	return r.preparedQuery.CreateExercise(ctx, entity.CreateExerciseParams{
		Subject:    subject,
		CategoryID: categoryId,
		UserID:     userId,
	})
}

func (r *repo) FetchByDateTime(ctx context.Context, userId uuid.UUID, year, month int) (resp []entity.FetchByDateTimeRow, err error) {
	start, end := timex.GetDateForAMonthToUnixMilliSecond(year, month)
	resp, err = r.preparedQuery.FetchByDateTime(ctx, entity.FetchByDateTimeParams{
		UserID:      userId,
		CreatedAt:   start,
		CreatedAt_2: end,
	})
	if err != nil {
		return
	}
	return
}

func (r *repo) GetTodayExerciseCount(ctx context.Context, userId uuid.UUID, time int64) (int64, error) {
	start, end := timex.GetDateForADayUnixMillisecond(time)
	return r.preparedQuery.GetTodayExerciseCount(ctx, entity.GetTodayExerciseCountParams{
		UserID:      userId,
		CreatedAt:   start,
		CreatedAt_2: end,
	})
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

// FetchTodayExerciseByUserId
// TODO: bypass fix
func (r *repo) FetchTodayExerciseByUserId(ctx context.Context, userId uuid.UUID, time int64) ([]entity.FetchTodayExerciseByUserIdRow, error) {
	start, end := timex.GetDateForADayUnixMillisecond(time)

	return r.preparedQuery.FetchTodayExerciseByUserId(ctx, entity.FetchTodayExerciseByUserIdParams{
		UserID:      userId,
		CreatedAt:   start,
		CreatedAt_2: end,
	})
}

func (r *repo) DeleteHealth(ctx context.Context, userId uuid.UUID, id int64, time int64) error {
	start, end := timex.GetDateForADayUnixMillisecond(time)

	return r.preparedQuery.DeleteHealth(ctx, entity.DeleteHealthParams{
		UserID:      userId,
		ExerciseID:  id,
		CreatedAt:   start,
		CreatedAt_2: end,
	})
}
