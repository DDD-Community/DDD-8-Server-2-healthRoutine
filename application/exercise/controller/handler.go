package controller

import (
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
		ExerciseId int64 `json:"-" xml:"-" param:"exerciseId"`
	}
	if err = c.ParamsParser(&binder); err != nil {
		return err
	}

	err = h.useCase.DeleteExerciseUseCase.Use(c.Context(), binder.ExerciseId, userId)
	switch {
	case err == nil:
		return c.SendStatus(http.StatusNoContent)
	case errors.Is(err, exercise.ErrNotMatchUserId):
		err = response.ErrNotMatchUserId
		logger.Error(err)
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
		Weight     int32 `json:"weight" xml:"-"`
		Reps       int32 `json:"reps" xml:"-"`
		Set        int32 `json:"set" xml:"-"`
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
		Reps:       binder.Reps,
		Set:        binder.Set,
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

	userId, err := middlewares.ExtractUserId(c)
	if err != nil {
		return response.ErrUnauthorized
	}

	var binder struct {
		CategoryId int64 `json:"-" xml:"-" query:"category"`
	}

	if err := c.QueryParser(&binder); err != nil {
		logger.Error(err)
		return response.ErrorResponse(c, err, nil)
	}

	resp, err := h.useCase.FetchExerciseByCategoryIdUseCase.Use(c.Context(), userId, binder.CategoryId)
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
	case err == nil:
		return c.Status(http.StatusOK).JSON(controller.NewResponseBody(http.StatusOK, fetchTodayExerciseResultToData(resp)))
	case userId == uuid.Nil:
		return c.Status(http.StatusNotFound).JSON(response.ErrNotFoundUser)
	}

	return response.ErrorResponse(c, err, func(err error) {
		logger.Error(err)
		logger.Error(binder.Time)
		logger.Error("failed to fetch today exercise")
	})
}

func (h *Handler) deleteHealth(c *fiber.Ctx) error {
	logger := h.log()
	defer logger.Sync()

	var binder struct {
		HealthId uuid.UUID `json:"-" xml:"-" param:"healthId"`
	}
	if err := c.ParamsParser(&binder); err != nil {
		return response.ErrorResponse(c, err, nil)
	}

	err := h.useCase.DeleteHealthUseCase.Use(c.Context(), binder.HealthId)
	if err != nil {
		return response.ErrorResponse(c, err, func(err error) {
			logger.Error("failed to delete health")
			logger.Error(err)
		})
	}

	return c.SendStatus(http.StatusNoContent)
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

func (h *Handler) getWater(c *fiber.Ctx) error {
	logger := h.log()
	defer logger.Sync()

	userId, err := middlewares.ExtractUserId(c)
	if err != nil {
		err = response.ErrUnauthorized
		return response.ErrorResponse(c, err, nil)
	}

	var getWater GetWaterData
	resp, err := h.useCase.GetWaterByUserIdUseCase.Use(c.Context(), userId)
	switch {
	case err == nil:
		return c.Status(http.StatusOK).JSON(controller.NewResponseBody(http.StatusOK, getWater.getWaterToGetWaterData(resp)))
	case errors.Is(err, user.ErrNoRecordDrink):
		return c.Status(http.StatusOK).JSON(controller.NewResponseBody(http.StatusOK, getWater.isZero()))
	}
	return response.ErrorResponse(c, err, func(err error) {
		logger.Error(err)
	})
}

func (h *Handler) createOrUpdateWater(c *fiber.Ctx) error {
	logger := h.log()
	defer logger.Sync()

	userId, err := middlewares.ExtractUserId(c)
	if err != nil {
		err = response.ErrUnauthorized
		return response.ErrorResponse(c, err, nil)
	}

	var binder struct {
		Capacity *int64 `json:"capacity" xml:"-" validate:"required,min=0"`
	}
	if err = c.BodyParser(&binder); err != nil {
		return err
	}

	validateErrors := util.ValidateStruct(&binder)
	if validateErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validateErrors)
	}

	err = h.useCase.CreateOrUpdateWaterUseCase.Use(c.Context(), userId, *binder.Capacity)
	if err != nil {
		return response.ErrorResponse(c, err, func(err error) {
			logger.Error(err)
		})
	}

	return c.Status(http.StatusCreated).JSON(controller.NewResponseBody(http.StatusCreated))
}
