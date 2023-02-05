package controller

import (
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
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

type Handler struct {
	useCase usecase.UserUseCases
}

func (*Handler) log() *zap.SugaredLogger {
	return log.Get().Named("USER_CONTROLLER")
}

func (h *Handler) signUp(c *fiber.Ctx) error {
	logger := h.log()
	defer logger.Sync()

	var binder struct {
		Nickname string `json:"nickname" xml:"-" validate:"required"`
		Email    string `json:"email" xml:"-" validate:"required"`
		Password string `json:"password" xml:"-" validate:"required"`
	}

	if err := c.BodyParser(&binder); err != nil {
		return err
	}

	if !util.CheckEmail(binder.Email) {
		err := response.ErrInvalidEmail
		logger.Error("failed to check email")
		logger.Error(err)
		return response.ErrorResponse(c, err, nil)
	}

	if !util.CheckPassword(binder.Password) {
		err := response.ErrInvalidPassword
		logger.Error("failed to check password")
		logger.Error(err)
		return response.ErrorResponse(c, err, nil)
	}

	resp, err := h.useCase.SignUpUseCase.Use(c.Context(), user.SignUpParams{
		Nickname: binder.Nickname,
		Password: binder.Password,
		Email:    binder.Email,
	})
	switch {
	case errors.Is(err, user.ErrEmailAlreadyExists):
		err = response.ErrEmailAlreadyExist
		return response.ErrorResponse(c, err, nil)
	case err != nil:
		return response.ErrorResponse(c, err, func(err error) {
			logger.Error("failed to sign up")
		})
	}

	return c.Status(http.StatusCreated).JSON(ResponseByHttpStatus(http.StatusCreated, domainModelToSignUpResp(resp)))
}

func (h *Handler) signIn(c *fiber.Ctx) error {
	logger := log.Get()
	defer logger.Sync()

	var binder struct {
		Email    string `json:"email" xml:"-" validate:"required"`
		Password string `json:"password" xml:"-" validate:"required"`
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
			logger.Error("failed to sign in")
		})
	}

	res := ResponseByHttpStatus(http.StatusOK, domainModelToSignInResp(resp, token))

	return c.Status(http.StatusOK).JSON(res)
}

func (h *Handler) checkEmailValidation(c *fiber.Ctx) error {
	logger := log.Get()
	defer logger.Sync()

	var binder struct {
		Email string `json:"email" xml:"-" validate:"required"`
	}
	validateErrors := util.ValidateStruct(&binder) // TODO: validator 분리 필요
	if err := c.BodyParser(&binder); err != nil {
		return response.ErrorResponse(c, err, nil)
	}
	if validateErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validateErrors)
	}

	err := h.useCase.EmailValidationUseCase.Use(c.Context(), binder.Email)
	switch {
	case errors.Is(err, user.ErrEmailAlreadyExists):
		err = response.ErrEmailAlreadyExist
		return response.ErrorResponse(c, err, nil)
	case err != nil:
		return response.ErrorResponse(c, err, func(err error) {
			logger.Error("failed to check email validation")
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
			logger.Error("failed to get profile")
		})
	}

	res := domainModelToProfileResp(resp)

	return c.Status(http.StatusOK).JSON(ResponseByHttpStatus(http.StatusOK, res))

}

func (h *Handler) updateProfile(c *fiber.Ctx) error {
	logger := log.Get()
	defer logger.Sync()

	userId, err := middlewares.ExtractUserId(c)
	if err != nil {
		return response.ErrUnauthorized
	}

	var binder struct {
		Nickname string `json:"nickname" xml:"-" validate:"required"`
	}
	if err = c.BodyParser(&binder); err != nil {
		logger.Error(err)
		return response.ErrorResponse(c, err, nil)
	}

	validateErrors := util.ValidateStruct(&binder)
	if validateErrors != nil {
		return c.Status(http.StatusBadRequest).JSON(validateErrors)
	}

	err = h.useCase.UpdateProfileUseCase.Use(c.Context(), user.UpdateProfileParams{
		Id:       userId,
		Nickname: binder.Nickname,
	})
	switch {
	case errors.Is(err, user.ErrNicknameAlreadyExists):
		err = response.ErrNicknameAlreadyExist
		return response.ErrorResponse(c, err, nil)
	case err != nil:
		return response.ErrorResponse(c, err, func(err error) {
			logger.Error("failed to update nickname")
		})
	}

	// http 표준은 204 No Content 가 맞지만
	// response body 로 처리하고있기에, 임시로 200 OK 로 처리 중
	return c.Status(http.StatusOK).JSON(NewResponseBody(http.StatusOK))

}

func (h *Handler) uploadProfileImg(c *fiber.Ctx) error {
	logger := log.Get()
	defer logger.Sync()

	userId, err := middlewares.ExtractUserId(c)
	if err != nil {
		return response.ErrUnauthorized
	}

	file, err := c.FormFile("file")
	if err != nil {
		logger.Error(err)
		return response.ErrorResponse(c, err, nil)
	}

	src, err := file.Open()
	if err != nil {
		logger.Error(err)
		return response.ErrorResponse(c, err, nil)
	}

	defer src.Close()

	if !strings.HasPrefix(file.Header["Content-Type"][0], "image/") {
		logger.Error("invalid content type")
		err = response.ErrInvalidContentType
		return response.ErrorResponse(c, err, nil)
	}

	url, err := h.useCase.UploadTemporaryProfileUseCase.Use(c.Context(), user.UploadTemporaryProfileParams{
		Id:            userId,
		Filename:      file.Filename,
		ContentType:   file.Header["Content-Type"][0],
		ContentLength: file.Size,
		ProfileImg:    src,
	})
	if err != nil {
		return response.ErrorResponse(c, err, func(err error) {
			logger.Error("failed to update profile image")
		})
	}

	var res struct {
		ProfileImageUrl string `json:"profileImageUrl"`
	}

	res.ProfileImageUrl = url

	// http 표준은 204 No Content 가 맞지만
	// response body 로 처리하고있기에, 임시로 200 OK 로 처리 중
	return c.Status(http.StatusOK).JSON(NewResponseBody(http.StatusOK, res))
}

func (h *Handler) getBadge(c *fiber.Ctx) error {
	logger := h.log()
	defer logger.Sync()

	userId, err := middlewares.ExtractUserId(c)
	if err != nil {
		return response.ErrUnauthorized
	}

	resp, err := h.useCase.GetBadgeUseCase.Use(c.Context(), userId)
	if err != nil {
		return response.ErrorResponse(c, err, func(err error) {
			logger.Error("failed to get badge")
			logger.Error(err)
		})
	}
	return c.Status(http.StatusOK).JSON(NewResponseBody(http.StatusOK, resp))
}
