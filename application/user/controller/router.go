package controller

import (
	"github.com/gofiber/fiber/v2"
	"healthRoutine/application/user/usecase"
)

func BindHandler(app *fiber.App, cases usecase.UseCases) {
	var handler = Handler{
		useCase: cases,
	}

	api := app.Group("api/v1")

	api.Post("/user/register", handler.signUp)
	api.Post("/user/login", handler.signIn)
	api.Post("/user/validate/email", handler.checkEmailValidation)
}
