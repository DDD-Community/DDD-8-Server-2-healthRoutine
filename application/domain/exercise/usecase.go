package exercise

import (
	"context"
	"github.com/google/uuid"
	entity "healthRoutine/pkgs/database/sqlc/exercise"
)

type CreateHistoryParams struct {
	UserId     uuid.UUID
	ExerciseId int64
	Weight     int32
	Reps       int32
	Set        int32
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
	Use(ctx context.Context, params CreateExerciseParams) (int64, error)
}

type FetchExerciseResult struct {
	Id       int64           `json:"id"`
	Subject  string          `json:"subject"`
	Exercise []ExerciseModel `json:"exercise"`
}

type FetchExerciseByCategoryIdUseCase interface {
	Use(ctx context.Context, userId uuid.UUID) ([]FetchExerciseResult, error)
}

type FetchCategoriesUseCase interface {
	Use(ctx context.Context) ([]ExerciseCategoryModel, error)
}

type FetchTodayExerciseResult struct {
	Id              uuid.UUID
	ExerciseSubject string
	CategorySubject string
	Weight          int32
	Set             int32
	Reps            int32
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
	Level         int32
	IsFutureDays  bool
}

type FetchByDatetimeUseCase interface {
	Use(ctx context.Context, userId uuid.UUID, year, month int) (FetchByDatetimeResult, error)
}

type DeleteExerciseUseCase interface {
	Use(ctx context.Context, id int64, userId uuid.UUID) error
}

type DeleteHealthUseCase interface {
	Use(ctx context.Context, id uuid.UUID) error
}

type GetWaterByUserIdUseCase interface {
	Use(ctx context.Context, userId uuid.UUID, y, m, d int) (entity.Water, error)
}

type GetTodayWaterByUserIdUseCase interface {
	Use(ctx context.Context, userId uuid.UUID) (entity.Water, error)
}

type CreateOrUpdateWaterUseCase interface {
	Use(ctx context.Context, userId uuid.UUID, capacity int64) error
}
