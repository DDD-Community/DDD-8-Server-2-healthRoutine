package main

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	controller2 "healthRoutine/application/exercise/controller"
	repository2 "healthRoutine/application/exercise/repository"
	usecase2 "healthRoutine/application/exercise/usecase"
	"healthRoutine/application/user/controller"
	"healthRoutine/application/user/repository"
	"healthRoutine/application/user/usecase"
	"healthRoutine/cmd"
	"healthRoutine/pkgs/database"
	"net/http"
)

const addr = ":3000"

func main() {
	app := fiber.New()
	db := database.Conn()

	// use fiber logger
	app.Use(logger.New())

	// if panic, recover
	app.Use(recover.New())

	app.All("/", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(http.StatusBadGateway)
	})

	userRepo := repository.NewUserRepository(db)
	exerciseRepo := repository2.NewExerciseRepository(db)
	defaultS3 := s3.NewFromConfig(cmd.GetAWSConfig())

	userUseCase := usecase.UserUseCases{
		SignUpUseCase:                 usecase.SignUpUseCase(userRepo),
		SignInUseCase:                 usecase.SignInUseCase(userRepo),
		EmailValidationUseCase:        usecase.EmailValidationUseCase(userRepo),
		GetProfileUseCase:             usecase.GetProfileUseCase(userRepo),
		UploadTemporaryProfileUseCase: usecase.UploadTemporaryProfileUseCase(userRepo, defaultS3),
		UpdateProfileUseCase:          usecase.UpdateProfileUseCase(userRepo, defaultS3),
	}

	exerciseUseCase := usecase2.ExerciseUseCase{
		CreateHistoryUseCase:              usecase2.CreateHistoryUseCase(exerciseRepo),
		CreateExerciseUseCase:             usecase2.CreateExerciseUseCase(exerciseRepo),
		FetchExerciseByCategoryIdUseCase:  usecase2.FetchExerciseByCategoryIdUseCase(exerciseRepo),
		FetchCategoriesUseCase:            usecase2.FetchCategoriesUseCase(exerciseRepo),
		FetchTodayExerciseByUserIdUseCase: usecase2.FetchTodayExerciseByUserIdUseCase(exerciseRepo),
		FetchByDatetimeUseCase:            usecase2.FetchByDatetimeUseCase(exerciseRepo, userRepo),
		DeleteExerciseUseCase:             usecase2.DeleteExerciseUseCase(exerciseRepo),
		DeleteHealthUseCase:               usecase2.DeleteHealthUseCase(exerciseRepo),
	}

	controller.BindUserHandler(app, userUseCase)
	controller2.BindExerciseHandler(app, exerciseUseCase)

	app.Listen(addr)
}
