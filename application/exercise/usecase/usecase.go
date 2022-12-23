package usecase

import "healthRoutine/application/domain/exercise"

type ExerciseUseCase struct {
	exercise.CreateExerciseUseCase
	exercise.FetchExerciseByCategoryIdUseCase
	exercise.FetchCategoriesUseCase
	exercise.FetchTodayExerciseByUserIdUseCase
}
