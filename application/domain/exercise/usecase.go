package exercise

import (
	"context"
	"github.com/google/uuid"
)

type CreateHistoryParams struct {
	UserId     uuid.UUID
	ExerciseId int64
	Weight     int32
	Set        int32
	Minute     int32
}

type CreateHistoryUseCase interface {
	Use(ctx context.Context, params CreateHistoryParams) error
}

type CreateExerciseParams struct {
	UserId     *uuid.UUID
	CategoryId int64
	Subject    string
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

type FetchTodayExerciseResult struct {
	ExerciseSubject string
	CategorySubject string
	Weight          int64
	Set             int64
	Count           int64
	CreatedAt       int64
}

type FetchTodayExerciseByUserIdUseCase interface {
	Use(ctx context.Context, userId uuid.UUID, time int64) ([]FetchTodayExerciseResult, error)
}

type FetchByDatetimeResult struct {
	Year           int
	Month          int
	WelcomeMessage string
	Data           []FetchByDatetimeDetail
}

type FetchByDatetimeDetail struct {
	Day           int
	TotalExercise int64
	Level         *int32
	IsFutureDays  bool
}

type FetchByDatetimeUseCase interface {
	Use(ctx context.Context, userId uuid.UUID, year, month int) (FetchByDatetimeResult, error)
}

type DeleteExerciseUseCase interface {
	Use(ctx context.Context, id int64, userId uuid.UUID) error
}

type DeleteHealthUseCase interface {
	Use(ctx context.Context, userId uuid.UUID, id int64, time int64) error
}
