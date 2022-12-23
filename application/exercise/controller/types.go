package controller

import (
	"healthRoutine/application/domain/exercise"
)

// TODO: refactor

type fetchExerciseData struct {
	Id      int64  `json:"id"`
	Subject string `json:"subject"`
}

func exerciseDomainToData(model []exercise.ExerciseModel) (res []fetchExerciseData) {
	res = make([]fetchExerciseData, 0, len(model))
	for _, v := range model {
		res = append(res, fetchExerciseData{
			Id:      v.Exercise.ID,
			Subject: v.Exercise.Subject,
		})
	}

	return
}

type fetchCategoriesData struct {
	Id      int64
	Subject string
}

func exerciseCategoryDomainToData(model []exercise.ExerciseCategoryModel) (res []fetchCategoriesData) {
	res = make([]fetchCategoriesData, 0, len(model))
	for _, v := range model {
		res = append(res, fetchCategoriesData{
			Id:      v.ExerciseCategory.ID,
			Subject: v.ExerciseCategory.Subject,
		})
	}

	return
}
