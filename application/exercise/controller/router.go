package controller

import (
	"github.com/gofiber/fiber/v2"
	"healthRoutine/application/exercise/usecase"
	"healthRoutine/pkgs/middlewares"
)

func BindExerciseHandler(app *fiber.App, cases usecase.ExerciseUseCase) {
	var handler = Handler{
		useCase: cases,
	}

	api := app.Group("api/v1")

	// exercise
	api.Post("/exercise", middlewares.AuthRequired(), handler.createExercise) // 운동 생성
	api.Get("/exercise", handler.fetchExerciseByCategoryId)                   // 카테고리 별 운동 리스트

	// exercise category
	api.Get("/exercise/category", handler.fetchCategories)
	api.Get("/exercise/today", middlewares.AuthRequired(), handler.fetchTodayExercise)
}
