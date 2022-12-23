package controller

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"healthRoutine/application/domain/exercise"
	"healthRoutine/application/exercise/usecase"
	"healthRoutine/application/user/controller"
	"healthRoutine/pkgs/errors/response"
	"healthRoutine/pkgs/log"
	"healthRoutine/pkgs/middlewares"
	"healthRoutine/pkgs/util"
	"net/http"
)

const (
	named = "EXERCISE_CONTROLLER"
)

type Handler struct {
	useCase usecase.ExerciseUseCase
}

func (h *Handler) log() *zap.SugaredLogger {
	return log.Get().Named(named)
}

func (h *Handler) createExercise(c *fiber.Ctx) error {
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

	errors := util.ValidateStruct(&binder)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	if err = c.BodyParser(&binder); err != nil {
		return err
	}

	err = h.useCase.CreateExerciseUseCase.Use(c.Context(), exercise.CreateExerciseParams{
		UserId:     userId,
		ExerciseId: binder.ExerciseId,
		Weight:     binder.Weight,
		Set:        binder.Set,
		Minute:     binder.Minute,
	})
	if err != nil {
		logger.Error(err)
		return response.ErrorResponse(c, err, nil)
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
	case err == sql.ErrNoRows:
		return c.Status(http.StatusNotFound).JSON(response.ErrNotFoundUser)
	case userId == uuid.Nil:
		return c.Status(http.StatusNotFound).JSON(response.ErrNotFoundUser)
	case err != nil:
		return response.ErrorResponse(c, err, func(err error) {
			logger.Error(err)
			logger.Error("failed to fetch today exercise")
		})
	}

	// TODO: type binding
	return c.Status(http.StatusOK).JSON(controller.NewResponseBody(http.StatusOK, resp))
}
