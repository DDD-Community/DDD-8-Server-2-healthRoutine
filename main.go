package main

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"healthRoutine/application/user/controller"
	"healthRoutine/application/user/repository"
	"healthRoutine/application/user/usecase"
	"healthRoutine/cmd"
	"healthRoutine/pkgs/database"
)

const addr = ":3000"

func main() {
	app := fiber.New()
	db := database.Conn()

	// use fiber logger
	app.Use(logger.New())

	userRepo := repository.NewUserRepository(db)
	defaultS3 := s3.NewFromConfig(cmd.GetAWSConfig())

	useCase := usecase.UseCases{
		SignUpUseCase:           usecase.SignUpUseCase(userRepo),
		SignInUseCase:           usecase.SignInUseCase(userRepo),
		EmailValidationUseCase:  usecase.EmailValidationUseCase(userRepo),
		GetProfileUseCase:       usecase.GetProfileUseCase(userRepo),
		UpdateProfileImgUseCase: usecase.UpdateProfileImgUseCase(userRepo, defaultS3),
		UpdateNicknameUseCase:   usecase.UpdateNicknameUseCase(userRepo),
	}

	controller.BindHandler(app, useCase)

	app.Listen(addr)

}
