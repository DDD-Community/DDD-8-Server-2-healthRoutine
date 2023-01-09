package exercise

import (
	"context"
	"github.com/google/uuid"
)

type CreateExerciseParams struct {
	UserId     uuid.UUID
	ExerciseId int64
	Weight     int32
	Set        int32
	Minute     int32
}

type CreateExerciseUseCase interface {
	Use(ctx context.Context, params CreateExerciseParams) error
}

type FetchExerciseByCategoryIdUseCase interface {
	Use(ctx context.Context, categoryId int64) ([]ExerciseModel, error)
}

type FetchCategoriesUseCase interface {
	Use(ctx context.Context) ([]ExerciseCategoryModel, error)
}

type FetchTodayExerciseParams struct {
	Subject string
	Count   int64
}

type FetchTodayExerciseByUserIdUseCase interface {
	Use(ctx context.Context, userId uuid.UUID, time int64) ([]FetchTodayExerciseParams, error)
}

type FetchByDatetimeResult struct {
	Year           int
	Month          int
	TotalOfMonth   int
	WelcomeMessage string
	Data           []FetchByDatetimeDetail
}

type FetchByDatetimeDetail struct {
	Day           int
	TotalExercise int64
	Level         int32
}

type FetchByDatetimeUseCase interface {
	Use(ctx context.Context, userId uuid.UUID, year, month int) (FetchByDatetimeResult, error)
}
