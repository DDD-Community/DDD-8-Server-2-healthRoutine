package main

import (
	"github.com/gofiber/fiber/v2"
	"healthRoutine/application/user/controller"
	"healthRoutine/application/user/repository"
	"healthRoutine/application/user/usecase"
	"healthRoutine/pkgs/database"
)

const addr = ":3000"

func main() {
	app := fiber.New()
	db := database.Conn()

	userRepo := repository.NewUserRepository(db)

	useCase := usecase.UseCases{
		SignUpUseCase: usecase.SignUpUseCase(userRepo),
		SignInUseCase: usecase.SignInUseCase(userRepo),
	}

	controller.BindHandler(app, useCase)

	app.Listen(addr)

}
