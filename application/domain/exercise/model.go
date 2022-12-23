package exercise

import entity "healthRoutine/pkgs/database/sqlc/exercise"

type DomainModel struct {
	Health entity.Health
}

type ExerciseModel struct {
	Exercise entity.Exercise
}

type ExerciseCategoryModel struct {
	ExerciseCategory entity.ExerciseCategory
}
