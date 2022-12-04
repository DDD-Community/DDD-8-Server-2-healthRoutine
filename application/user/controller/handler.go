package controller

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"healthRoutine/application/domain/user"
	"healthRoutine/application/user/usecase"
	"healthRoutine/pkgs/errors/response"
	"healthRoutine/pkgs/util"
	"log"
	"net/http"
)

type Handler struct {
	useCase usecase.UseCases
}

func (h *Handler) signUp(c *fiber.Ctx) error {
	var binder struct {
		Nickname string `json:"nickname" xml:"-"`
		Email    string `json:"email" xml:"-"`
		Password string `json:"password" xml:"-"`
	}
	if err := c.BodyParser(&binder); err != nil {
		return err
	}

	// TODO: need to apply logger
	if !util.CheckEmailRegex(binder.Email) {
		err := response.ErrInvalidEmail
		return response.ErrorResponse(c, err, nil)
	}

	if !util.CheckPassword(binder.Password) {
		err := response.ErrInvalidPassword
		return response.ErrorResponse(c, err, nil)
	}

	err := h.useCase.SignUpUseCase.Use(c.Context(), user.SignUpParams{
		Nickname: binder.Nickname,
		Password: binder.Password,
		Email:    binder.Email,
	})
	if err != nil {
		return response.ErrorResponse(c, err, func(err error) {
			log.Print("failed to sign up")
		})
	}

	return c.Status(http.StatusCreated).Send(nil)
}

func (h *Handler) signIn(c *fiber.Ctx) error {
	var binder struct {
		Email    string `json:"email" xml:"-"`
		Password string `json:"password" xml:"-"`
	}
	if err := c.BodyParser(&binder); err != nil {
		return response.ErrorResponse(c, err, nil)
	}

	resp, err := h.useCase.SignInUseCase.Use(c.Context(), binder.Email, binder.Password)
	switch {
	case err == sql.ErrNoRows:
		err = response.ErrNotFoundUser
		return response.ErrorResponse(c, err, nil)
	case err == bcrypt.ErrMismatchedHashAndPassword:
		err = response.ErrWrongPassword
		return response.ErrorResponse(c, err, nil)
	case err != nil:
		return response.ErrorResponse(c, err, func(err error) {
			log.Print("failed to sign in")
		})
	}

	return c.Status(http.StatusOK).JSON(map[string]string{
		"token": resp,
	})
}
