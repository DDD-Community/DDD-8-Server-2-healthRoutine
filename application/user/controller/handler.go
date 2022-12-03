package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"healthRoutine/application/domain/user"
	"healthRoutine/application/user/usecase"
	"healthRoutine/pkgs/util"
	"net/http"
)

type Handler struct {
	useCase usecase.UseCases
}

// TODO: open api 3.0
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
		fmt.Println("not match regex")
	}

	if !util.CheckPassword(binder.Password) {
		fmt.Println("not match regex")
	}

	err := h.useCase.SignUpUseCase.Use(c.Context(), user.SignUpParams{
		Nickname: binder.Nickname,
		Password: binder.Password,
		Email:    binder.Email,
	})
	if err != nil {
		panic(err) // not panic need return
	}

	return c.Status(http.StatusCreated).Send(nil)
}

func (h *Handler) signIn(c *fiber.Ctx) error {
	var binder struct {
		Email    string `json:"email" xml:"-"`
		Password string `json:"password" xml:"-"`
	}
	if err := c.BodyParser(&binder); err != nil {
		return err
	}

	resp, err := h.useCase.SignInUseCase.Use(c.Context(), binder.Email, binder.Password)
	if err != nil {
		fmt.Println("resp err")
	}

	return c.Status(http.StatusOK).JSON(map[string]string{
		"token": resp,
	})
}
