package response

import (
	"net/http"
)

var (
	ErrUnauthorized         = NewErrorResponse(http.StatusUnauthorized, "unauthorized")
	ErrInvalidEmail         = NewErrorResponse(http.StatusBadRequest, "invalid email")
	ErrInvalidPassword      = NewErrorResponse(http.StatusBadRequest, "invalid password")
	ErrNotFoundUser         = NewErrorResponse(http.StatusNotFound, "not found user")
	ErrWrongPassword        = NewErrorResponse(http.StatusUnauthorized, "wrong password")
	ErrEmailAlreadyExist    = NewErrorResponse(http.StatusConflict, "email already exists")
	ErrNicknameAlreadyExist = NewErrorResponse(http.StatusConflict, "nickname already exists")
	ErrInvalidContentType   = NewErrorResponse(http.StatusUnsupportedMediaType, "invalid content type")
	ErrNoBadge              = NewErrorResponse(http.StatusNotFound, "don't have any badge")

	ErrCharacterLimit = NewErrorResponse(http.StatusBadRequest, "20 character limit error")
	ErrNotMatchUserId = NewErrorResponse(http.StatusBadRequest, "user id doesn't match")
)
