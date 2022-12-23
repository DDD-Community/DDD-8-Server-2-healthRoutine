package controller

import (
	"github.com/gofiber/fiber/v2"
	"healthRoutine/application/user/usecase"
	"healthRoutine/pkgs/middlewares"
)

func BindHandler(app *fiber.App, cases usecase.UseCases) {
	var handler = Handler{
		useCase: cases,
	}

	api := app.Group("api/v1")

	api.Post("/user/register", handler.signUp)
	api.Post("/user/login", handler.signIn)
	api.Post("/user/validate/email", handler.checkEmailValidation)
	api.Get("/user/profile", middlewares.AuthRequired(), handler.getProfile)
	api.Put("/user/profile/image", middlewares.AuthRequired(), handler.updateProfileImg)
	api.Put("/user/profile/nickname", middlewares.AuthRequired(), handler.updateNickname)
}
