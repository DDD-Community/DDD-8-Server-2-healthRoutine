package exercise

import entity "healthRoutine/pkgs/database/sqlc/exercise"

type DomainModel struct {
	Health entity.Health
}

type ExerciseModel struct {
	Id      int64  `json:"id"`
	Subject string `json:"subject"`
}

type ExerciseCategoryModel struct {
	ExerciseCategory entity.ExerciseCategory
}
