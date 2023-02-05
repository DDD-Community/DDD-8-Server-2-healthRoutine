package controller

import (
	"github.com/gofiber/fiber/v2"
	"healthRoutine/application/user/usecase"
	"healthRoutine/pkgs/middlewares"
)

func BindUserHandler(app *fiber.App, cases usecase.UserUseCases) {
	var handler = Handler{
		useCase: cases,
	}

	api := app.Group("/api/v1")

	api.Post("/user/register", handler.signUp)
	api.Post("/user/login", handler.signIn)
	api.Post("/user/validate/email", handler.checkEmailValidation)
	api.Get("/user/profile", middlewares.AuthRequired(), handler.getProfile)
	api.Put("/user/profile", middlewares.AuthRequired(), handler.updateProfile)
	api.Put("/user/profile/img-upload", middlewares.AuthRequired(), handler.uploadProfileImg)

	api.Get("/user/badge", middlewares.AuthRequired(), handler.getBadge)
}
