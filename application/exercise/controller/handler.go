package controller

import (
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"healthRoutine/application/domain/exercise"
	"healthRoutine/application/domain/user"
	"healthRoutine/application/exercise/usecase"
	"healthRoutine/application/user/controller"
	"healthRoutine/pkgs/errors/response"
	"healthRoutine/pkgs/log"
	"healthRoutine/pkgs/middlewares"
	"healthRoutine/pkgs/util"
	"net/http"
)

type Handler struct {
	useCase usecase.ExerciseUseCase
}

func (h *Handler) log() *zap.SugaredLogger {
	return log.Get().Named("EXERCISE_CONTROLLER")
}

func (h *Handler) deleteExercise(c *fiber.Ctx) error {
	logger := h.log()
	defer logger.Sync()

	userId, err := middlewares.ExtractUserId(c)
	if err != nil {
		return response.ErrUnauthorized
	}

	var binder struct {
		ExerciseId int64 `json:"exerciseId" xml:"-" validate:"required"`
	}
	if err = c.BodyParser(&binder); err != nil {
		return err
	}

	validatorErrors := util.ValidateStruct(&binder)
	if validatorErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validatorErrors)
	}

	err = h.useCase.DeleteExerciseUseCase.Use(c.Context(), binder.ExerciseId, userId)
	logger.Info(err)
	switch {
	case err == nil:
		return c.Status(http.StatusNoContent).JSON(controller.NewResponseBody(http.StatusNoContent))
	case errors.Is(err, exercise.ErrNotMatchUserId):
		err = response.ErrNotMatchUserId
		return response.ErrorResponse(c, err, nil)
	}

	return response.ErrorResponse(c, err, func(err error) {
		logger.Error("failed to delete exercise")
		logger.Error(err)
	})
}

func (h *Handler) createExercise(c *fiber.Ctx) error {
	logger := h.log()
	defer logger.Sync()

	userId, err := middlewares.ExtractUserId(c)
	if err != nil {
		return response.ErrUnauthorized
	}

	var binder struct {
		CategoryId int64  `json:"categoryId" xml:"-" validate:"required"`
		Subject    string `json:"subject" xml:"-" validate:"required"`
	}
	if err = c.BodyParser(&binder); err != nil {
		return err
	}

	validatorErrors := util.ValidateStruct(&binder)
	if validatorErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validatorErrors)
	}

	// 운동 20글자 이상이면 에러
	if len([]rune(binder.Subject)) > 20 {
		err = response.ErrCharacterLimit
		return response.ErrorResponse(c, err, nil)
	}

	in := exercise.CreateExerciseParams{
		UserId:     &userId,
		CategoryId: binder.CategoryId,
		Subject:    binder.Subject,
	}

	err = h.useCase.CreateExerciseUseCase.Use(c.Context(), in)
	if err != nil {
		return response.ErrorResponse(c, err, func(err error) {
			logger.Error("in", in)
			logger.Error(err)
		})
	}

	return c.Status(http.StatusCreated).JSON(controller.NewResponseBody(http.StatusCreated))
}

func (h *Handler) createExerciseHistory(c *fiber.Ctx) error {
	logger := h.log()
	defer logger.Sync()

	userId, err := middlewares.ExtractUserId(c)
	if err != nil {
		return response.ErrUnauthorized
	}

	var binder struct {
		ExerciseId int64 `json:"exerciseId" xml:"-" validate:"required"`
		Weight     int32 `json:"weight" xml:"-" validate:"required"`
		Set        int32 `json:"set" xml:"-" validate:"required"`
		Minute     int32 `json:"minute" xml:"-" validate:"required"`
	}
	if err = c.BodyParser(&binder); err != nil {
		return err
	}

	validateErrors := util.ValidateStruct(&binder)
	if validateErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validateErrors)
	}

	err = h.useCase.CreateHistoryUseCase.Use(c.Context(), exercise.CreateHistoryParams{
		UserId:     userId,
		ExerciseId: binder.ExerciseId,
		Weight:     binder.Weight,
		Set:        binder.Set,
		Minute:     binder.Minute,
	})
	if err != nil {
		return response.ErrorResponse(c, err, func(err error) {
			logger.Error(err)
		})
	}

	return c.Status(http.StatusCreated).JSON(controller.NewResponseBody(http.StatusCreated))
}

func (h *Handler) fetchExerciseByCategoryId(c *fiber.Ctx) error {
	logger := h.log()
	defer logger.Sync()

	var binder struct {
		CategoryId int64 `json:"-" xml:"-" query:"category"`
	}

	if err := c.QueryParser(&binder); err != nil {
		logger.Error(err)
		return response.ErrorResponse(c, err, nil)
	}

	resp, err := h.useCase.FetchExerciseByCategoryIdUseCase.Use(c.Context(), binder.CategoryId)
	if err != nil {
		return response.ErrorResponse(c, err, func(err error) {
			logger.Error(err)
			logger.Error("failed to fetch exercise list")
		})
	}

	res := exerciseDomainToData(resp)

	return c.Status(http.StatusOK).JSON(controller.NewResponseBody(http.StatusOK, res))
}

func (h *Handler) fetchCategories(c *fiber.Ctx) error {
	logger := h.log()
	defer logger.Sync()

	resp, err := h.useCase.FetchCategoriesUseCase.Use(c.Context())
	if err != nil {
		return response.ErrorResponse(c, err, func(err error) {
			logger.Error(err)
			logger.Error("failed to fetch categories")
		})
	}

	res := exerciseCategoryDomainToData(resp)

	return c.Status(http.StatusOK).JSON(controller.NewResponseBody(http.StatusOK, res))
}

func (h *Handler) fetchTodayExercise(c *fiber.Ctx) error {
	logger := h.log()
	defer logger.Sync()

	var binder struct {
		Time int64 `json:"-" xml:"-" query:"time"`
	}

	if err := c.QueryParser(&binder); err != nil {
		logger.Error(err)
		return response.ErrorResponse(c, err, nil)
	}

	userId, err := middlewares.ExtractUserId(c)
	if err != nil {
		return response.ErrUnauthorized
	}

	resp, err := h.useCase.FetchTodayExerciseByUserIdUseCase.Use(c.Context(), userId, binder.Time)
	switch {
	//TODO: sql error need to go repo
	case err == sql.ErrNoRows:
		return c.Status(http.StatusNotFound).JSON(response.ErrNotFoundUser)
	case userId == uuid.Nil:
		return c.Status(http.StatusNotFound).JSON(response.ErrNotFoundUser)
	case err != nil:
		return response.ErrorResponse(c, err, func(err error) {
			logger.Error(err)
			logger.Error(binder.Time)
			logger.Error("failed to fetch today exercise")
		})
	}

	// TODO: type binding
	return c.Status(http.StatusOK).JSON(controller.NewResponseBody(http.StatusOK, fetchTodayExerciseResultToData(resp)))
}

func (h *Handler) deleteHealth(c *fiber.Ctx) error {
	logger := h.log()
	defer logger.Sync()

	userId, err := middlewares.ExtractUserId(c)
	if err != nil {
		return response.ErrUnauthorized
	}

	var binder struct {
		ExerciseId int64 `json:"exerciseId" xml:"-" validate:"required"`
		CreatedAt  int64 `json:"createdAt" xml:"-" validate:"required"`
	}
	if err = c.BodyParser(&binder); err != nil {
		return response.ErrorResponse(c, err, nil)
	}

	validateErrors := util.ValidateStruct(&binder)
	if validateErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validateErrors)
	}

	err = h.useCase.DeleteHealthUseCase.Use(c.Context(), userId, binder.ExerciseId, binder.CreatedAt)
	if err != nil {
		return response.ErrorResponse(c, err, func(err error) {
			logger.Error("failed to delete health")
			logger.Error(err)
		})
	}

	return c.Status(http.StatusNoContent).JSON(controller.NewResponseBody(http.StatusNoContent))
}

func (h *Handler) fetchMonthly(c *fiber.Ctx) error {
	logger := h.log()
	defer logger.Sync()

	var binder struct {
		Year  int `json:"-" xml:"-" query:"year"`
		Month int `json:"-" xml:"-" query:"month"`
	}
	if err := c.QueryParser(&binder); err != nil {
		logger.Error(err)
		return response.ErrorResponse(c, err, nil)
	}

	userId, err := middlewares.ExtractUserId(c)
	if err != nil {
		err = response.ErrUnauthorized
		return response.ErrorResponse(c, err, nil)
	}

	resp, err := h.useCase.FetchByDatetimeUseCase.Use(c.Context(), userId, binder.Year, binder.Month)
	switch {
	case err == nil:
		return c.Status(http.StatusOK).JSON(controller.NewResponseBody(http.StatusOK, fetchByDatetimeResultTypes(resp)))
	case errors.Is(err, user.ErrUserNotFound):
		err = response.ErrNotFoundUser
		return response.ErrorResponse(c, err, nil)
	}

	return response.ErrorResponse(c, err, func(err error) {
		logger.Error(err)
	})
}
