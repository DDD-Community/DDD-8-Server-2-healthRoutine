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

	api := app.Group("/api/v1")

	// Need auth
	api.Get("/exercise/monthly", middlewares.AuthRequired(), handler.fetchMonthly)
	api.Get("/exercise", middlewares.AuthRequired(), handler.fetchExerciseByCategoryId)
	api.Post("/exercise", middlewares.AuthRequired(), handler.createExercise)
	api.Delete("/exercise/:exerciseId", middlewares.AuthRequired(), handler.deleteExercise)
	api.Get("/exercise/today", middlewares.AuthRequired(), handler.fetchTodayExercise)
	api.Delete("/exercise/today/:healthId", middlewares.AuthRequired(), handler.deleteHealth)
	api.Post("/exercise/history", middlewares.AuthRequired(), handler.createExerciseHistory)

	api.Get("/exercise/category", handler.fetchCategories)

	api.Get("/water/today", middlewares.AuthRequired(), handler.getTodayDrinkWater)
	api.Get("/water", middlewares.AuthRequired(), handler.getDrinkWaterByDate)
	api.Post("/water", middlewares.AuthRequired(), handler.createOrUpdateWater)
}
