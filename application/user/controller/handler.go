package controller

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"healthRoutine/application/domain/user"
	"healthRoutine/application/user/usecase"
	"healthRoutine/pkgs/errors/response"
	"healthRoutine/pkgs/log"
	"healthRoutine/pkgs/util"
	"net/http"
)

const (
	named = "USER_CONTROLLER"
)

type Handler struct {
	useCase usecase.UseCases
}

func (h *Handler) signUp(c *fiber.Ctx) error {
	logger := log.Get()
	defer logger.Sync()

	type Result struct {
		Email    string `json:"email"`
		Nickname string `json:"nickname"`
	}

	var binder struct {
		Nickname string `json:"nickname" xml:"-"`
		Email    string `json:"email" xml:"-"`
		Password string `json:"password" xml:"-"`
	}
	if err := c.BodyParser(&binder); err != nil {
		return err
	}

	if !util.CheckEmail(binder.Email) {
		err := response.ErrInvalidEmail
		logger.Named(named).Error("failed to check email")
		logger.Named(named).Error(err)
		return response.ErrorResponse(c, err, nil)
	}

	if !util.CheckPassword(binder.Password) {
		err := response.ErrInvalidPassword
		logger.Named(named).Error("failed to check password")
		logger.Named(named).Error(err)
		return response.ErrorResponse(c, err, nil)
	}

	resp, err := h.useCase.SignUpUseCase.Use(c.Context(), user.SignUpParams{
		Nickname: binder.Nickname,
		Password: binder.Password,
		Email:    binder.Email,
	})
	switch {
	case err == user.ErrEmailAlreadyExists:
		err = response.ErrEmailAlreadyExist
		return response.ErrorResponse(c, err, nil)
	case err == user.ErrNicknameAlreadyExists:
		err = response.ErrNicknameAlreadyExist
		return response.ErrorResponse(c, err, nil)
	case err != nil:
		return response.ErrorResponse(c, err, func(err error) {
			logger.Named(named).Error("failed to sign up")
		})
	}

	result := Result{
		Email:    resp.Email,
		Nickname: resp.Nickname,
	}

	return c.Status(http.StatusCreated).JSON(map[string]interface{}{
		"status":  "created",
		"code":    http.StatusCreated,
		"message": "successfully created",
		"result":  result,
	})
}

func (h *Handler) signIn(c *fiber.Ctx) error {
	logger := log.Get()
	defer logger.Sync()

	type Result struct {
		Email    string `json:"email"`
		Nickname string `json:"nickname"`
	}

	var binder struct {
		Email    string `json:"email" xml:"-"`
		Password string `json:"password" xml:"-"`
	}
	if err := c.BodyParser(&binder); err != nil {
		return response.ErrorResponse(c, err, nil)
	}

	resp, token, err := h.useCase.SignInUseCase.Use(c.Context(), binder.Email, binder.Password)
	switch {
	case err == sql.ErrNoRows:
		err = response.ErrNotFoundUser
		return response.ErrorResponse(c, err, nil)
	case err == bcrypt.ErrMismatchedHashAndPassword:
		err = response.ErrWrongPassword
		return response.ErrorResponse(c, err, nil)
	case err != nil:
		return response.ErrorResponse(c, err, func(err error) {
			logger.Named(named).Error("failed to sign in")
		})
	}

	result := Result{
		Email:    resp.Email,
		Nickname: resp.Nickname,
	}

	return c.Status(http.StatusOK).JSON(map[string]interface{}{
		"status":  "ok",
		"code":    http.StatusOK,
		"message": "success",
		"token":   token,
		"result":  result,
	})
}

func (h *Handler) checkEmailValidation(c *fiber.Ctx) error {
	logger := log.Get()
	defer logger.Sync()

	var binder struct {
		Email string `json:"email" xml:"-"`
	}
	if err := c.BodyParser(&binder); err != nil {
		return response.ErrorResponse(c, err, nil)
	}

	err := h.useCase.EmailValidationUseCase.Use(c.Context(), binder.Email)
	switch {
	case err == user.ErrEmailAlreadyExists:
		err = response.ErrEmailAlreadyExist
		return response.ErrorResponse(c, err, nil)
	case err != nil:
		return response.ErrorResponse(c, err, func(err error) {
			logger.Named(named).Error("failed to check email validation")
		})
	}

	return c.Status(http.StatusOK).JSON(map[string]interface{}{
		"status":  "ok",
		"code":    http.StatusOK,
		"message": "success",
	})
}
