package repository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"healthRoutine/application/domain/exercise"
	"healthRoutine/application/domain/user"
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
	weight, reps, set int32) error {
	return r.preparedQuery.Create(ctx, entity.CreateParams{
		ID:         uuid.New(),
		UserID:     userId,
		ExerciseID: exerciseId,
		Weight:     weight,
		Reps:       reps,
		Set:        set,
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

func (r *repo) CreateExercise(ctx context.Context, userId *uuid.UUID, categoryId int64, subject string) (int64, error) {
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

func (r *repo) FetchExerciseByCategoryId(ctx context.Context, userId uuid.UUID, categoryId int64) (res []exercise.ExerciseModel, err error) {
	resp, err := r.preparedQuery.FetchExerciseByCategoryId(ctx, entity.FetchExerciseByCategoryIdParams{
		CategoryID: categoryId,
		UserID:     &userId,
	})
	if err != nil {
		return
	}

	res = make([]exercise.ExerciseModel, 0, len(resp))
	for _, v := range resp {
		res = append(res, exercise.ExerciseModel{
			Id:      v.ID,
			Subject: v.Subject,
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

func (r *repo) DeleteHealth(ctx context.Context, id uuid.UUID) error {
	return r.preparedQuery.DeleteHealth(ctx, id)
}

// GetTodayWaterByUserId
// TODO: fix bypass
func (r *repo) GetTodayWaterByUserId(ctx context.Context, userId uuid.UUID) (resp entity.Water, err error) {
	start, end := timex.GetDateForADayUnixMillisecond(time.Now().UnixMilli())
	resp, err = r.preparedQuery.GetWaterByUserId(ctx, entity.GetWaterByUserIdParams{
		UserID:      userId,
		CreatedAt:   start,
		CreatedAt_2: end,
	})
	switch {
	case err == sql.ErrNoRows:
		err = user.ErrNoRecordDrink
		return
	case err != nil:
		return
	}

	return
}

// GetWaterByUserId
// TODO: fix bypass
func (r *repo) GetWaterByUserId(ctx context.Context, userId uuid.UUID, y, m, d int) (resp entity.Water, err error) {
	t := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local)
	start, end := timex.GetDateForADayUnixMillisecond(t.UnixMilli())
	resp, err = r.preparedQuery.GetWaterByUserId(ctx, entity.GetWaterByUserIdParams{
		UserID:      userId,
		CreatedAt:   start,
		CreatedAt_2: end,
	})
	switch {
	case err == sql.ErrNoRows:
		err = user.ErrNoRecordDrink
		return
	case err != nil:
		return
	}

	return
}

func (r *repo) CreateOrUpdateWater(ctx context.Context, userId uuid.UUID, capacity int64) error {
	now := time.Now().UnixMilli()
	date := time.UnixMilli(now).Format("2006-01-02")

	return r.preparedQuery.CreateOrUpdateWater(ctx, entity.CreateOrUpdateWaterParams{
		UserID:      userId,
		Capacity:    capacity,
		Unit:        "ML",
		Date:        date,
		CreatedAt:   now,
		UpdatedAt:   now,
		Capacity_2:  capacity,
		UpdatedAt_2: now,
	})
}

func (r *repo) CountExerciseHistoryByUserId(ctx context.Context, userId uuid.UUID) (int64, error) {
	return r.preparedQuery.CountExerciseHistoryByUserId(ctx, userId)
}

func (r *repo) CountDrinkHistoryByUserId(ctx context.Context, userId uuid.UUID) (int64, error) {
	return r.preparedQuery.CountDrinkHistoryByUserId(ctx, userId)
}
