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
	Id      int64  `json:"id"`
	Subject string `json:"subject"`
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

type fetchTodayExerciseData struct {
	ExerciseId      int64  `json:"exerciseId"`
	ExerciseSubject string `json:"exerciseSubject"`
	CategorySubject string `json:"categorySubject"`
	Weight          int64  `json:"weight"`
	Set             int64  `json:"set"`
	Count           int64  `json:"count"`
	CreatedAt       int64  `json:"createdAt"`
}

func fetchTodayExerciseResultToData(list []exercise.FetchTodayExerciseResult) (res []fetchTodayExerciseData) {
	res = make([]fetchTodayExerciseData, 0, len(list))
	for _, v := range list {
		res = append(res, fetchTodayExerciseData{
			ExerciseId:      v.ExerciseId,
			ExerciseSubject: v.ExerciseSubject,
			CategorySubject: v.CategorySubject,
			Weight:          v.Weight,
			Set:             v.Set,
			Count:           v.Count,
			CreatedAt:       v.CreatedAt,
		})
	}

	return
}

type fetchByDatetimeResult struct {
	Year           int                     `json:"year"`
	Month          int                     `json:"month"`
	WelcomeMessage string                  `json:"welcomeMessage"`
	Data           []fetchByDatetimeDetail `json:"data"`
}

type fetchByDatetimeDetail struct {
	Day           int   `json:"day"`
	TotalExercise int64 `json:"totalExercise"`
	Level         int32 `json:"level"`
	IsFutureDays  bool  `json:"isFutureDays"`
}

func fetchByDatetimeResultTypes(result exercise.FetchByDatetimeResult) fetchByDatetimeResult {
	data := make([]fetchByDatetimeDetail, 0, len(result.Data))
	for _, v := range result.Data {
		data = append(data, fetchByDatetimeDetail{
			Day:           v.Day,
			TotalExercise: v.TotalExercise,
			Level:         v.Level,
			IsFutureDays:  v.IsFutureDays,
		})
	}

	return fetchByDatetimeResult{
		Year:           result.Year,
		Month:          result.Month,
		WelcomeMessage: result.WelcomeMessage,
		Data:           data,
	}

}
