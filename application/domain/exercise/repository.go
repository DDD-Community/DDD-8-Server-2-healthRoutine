package exercise

import (
	"context"
	"github.com/google/uuid"
	entity "healthRoutine/pkgs/database/sqlc/exercise"
)

type CreateParams struct {
	UserId     uuid.UUID
	ExerciseId int64
	Weight     int32
	Set        int32
	Minute     int32
}

type Repository interface {
	Create(
		ctx context.Context,
		userId uuid.UUID,
		exerciseId int64,
		weight, set, minute int32) error
	GetExerciseById(ctx context.Context, id int64) (entity.Exercise, error)
	CreateExercise(ctx context.Context, userId *uuid.UUID, categoryId int64, subject string) error
	DeleteExercise(ctx context.Context, id int64, userId uuid.UUID) error
	FetchExerciseByCategoryId(ctx context.Context, categoryId int64) ([]ExerciseModel, error)
	FetchByDateTime(ctx context.Context, userId uuid.UUID, year, month int) (resp []entity.FetchByDateTimeRow, err error)
	GetTodayExerciseCount(ctx context.Context, userId uuid.UUID, time int64) (int64, error)
	FetchCategories(ctx context.Context) ([]ExerciseCategoryModel, error)
	FetchTodayExerciseByUserId(ctx context.Context, userId uuid.UUID, time int64) ([]entity.FetchTodayExerciseByUserIdRow, error)
	DeleteHealth(ctx context.Context, userId uuid.UUID, id int64, time int64) error
}
