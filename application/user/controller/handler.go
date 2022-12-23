package controller

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"healthRoutine/application/domain/user"
	"healthRoutine/application/user/usecase"
	"healthRoutine/pkgs/errors/response"
	"healthRoutine/pkgs/log"
	"healthRoutine/pkgs/middlewares"
	"healthRoutine/pkgs/util"
	"net/http"
	"strings"
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

	return c.Status(http.StatusCreated).JSON(ResponseByHttpStatus(http.StatusCreated, domainModelToSignUpResp(resp)))
}

func (h *Handler) signIn(c *fiber.Ctx) error {
	logger := log.Get()
	defer logger.Sync()

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

	res := ResponseByHttpStatus(http.StatusOK, domainModelToSignInResp(resp, token))

	return c.Status(http.StatusOK).JSON(res)
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

	return c.Status(http.StatusOK).JSON(ResponseByHttpStatus(http.StatusOK))
}

func (h *Handler) getProfile(c *fiber.Ctx) error {
	logger := log.Get()
	defer logger.Sync()

	userId, err := middlewares.ExtractUserId(c)
	if err != nil {
		return response.ErrUnauthorized
	}

	resp, err := h.useCase.GetProfileUseCase.Use(c.Context(), userId)
	switch {
	case err == sql.ErrNoRows:
		err = response.ErrNotFoundUser
		return response.ErrorResponse(c, err, nil)
	case err != nil:
		return response.ErrorResponse(c, err, func(err error) {
			logger.Named(named).Error("failed to get profile")
		})
	}

	res := domainModelToProfileResp(resp)

	return c.Status(http.StatusOK).JSON(ResponseByHttpStatus(http.StatusOK, res))

}

func (h *Handler) updateNickname(c *fiber.Ctx) error {
	logger := log.Get()
	defer logger.Sync()

	userId, err := middlewares.ExtractUserId(c)
	if err != nil {
		return response.ErrUnauthorized
	}

	var binder struct {
		Nickname string `json:"nickname" xml:"-"`
	}
	if err = c.BodyParser(&binder); err != nil {
		logger.Named(named).Error(err)
		return response.ErrorResponse(c, err, nil)
	}

	err = h.useCase.UpdateNicknameUseCase.Use(c.Context(), userId, binder.Nickname)
	if err != nil {
		return response.ErrorResponse(c, err, func(err error) {
			logger.Named(named).Error("failed to update nickname")
		})
	}

	return c.SendStatus(http.StatusNoContent)

}

func (h *Handler) updateProfileImg(c *fiber.Ctx) error {
	logger := log.Get()
	defer logger.Sync()

	userId, err := middlewares.ExtractUserId(c)
	if err != nil {
		return response.ErrUnauthorized
	}

	file, err := c.FormFile("file")
	if err != nil {
		logger.Named(named).Error(err)
		return response.ErrorResponse(c, err, nil)
	}

	src, err := file.Open()
	if err != nil {
		logger.Named(named).Error(err)
		return response.ErrorResponse(c, err, nil)
	}

	defer src.Close()
	buffer := make([]byte, 512)
	_, err = src.Read(buffer)
	if err != nil {
		logger.Named(named).Error(err)
		return response.ErrorResponse(c, err, nil)
	}

	contentType := http.DetectContentType(buffer)
	if !strings.HasPrefix(contentType, "image/") {
		logger.Error("invalid content type")
		err = response.ErrInvalidContentType
		return response.ErrorResponse(c, err, nil)
	}

	err = h.useCase.UpdateProfileImgUseCase.Use(c.Context(), user.UpdateProfileImgParams{
		Id:         userId,
		Filename:   file.Filename,
		ProfileImg: src,
	})
	if err != nil {
		return response.ErrorResponse(c, err, func(err error) {
			logger.Named(named).Error("failed to update profile image")
		})
	}

	return c.SendStatus(http.StatusNoContent)
}
