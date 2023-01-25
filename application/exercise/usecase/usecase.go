package usecase

import "healthRoutine/application/domain/exercise"

type ExerciseUseCase struct {
	exercise.CreateHistoryUseCase
	exercise.CreateExerciseUseCase
	exercise.FetchExerciseByCategoryIdUseCase
	exercise.FetchCategoriesUseCase
	exercise.FetchTodayExerciseByUserIdUseCase
	exercise.FetchByDatetimeUseCase
	exercise.DeleteExerciseUseCase
	exercise.DeleteHealthUseCase
	exercise.GetWaterByUserIdUseCase
	exercise.CreateOrUpdateWaterUseCase
}
