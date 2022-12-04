package response

import "net/http"

var (
	ErrUnauthorized    = NewErrorResponse(http.StatusUnauthorized, "unauthorized")
	ErrInvalidEmail    = NewErrorResponse(http.StatusBadRequest, "invalid email")
	ErrInvalidPassword = NewErrorResponse(http.StatusBadRequest, "invalid password")
	ErrNotFoundUser    = NewErrorResponse(http.StatusNotFound, "not found user")
	ErrWrongPassword   = NewErrorResponse(http.StatusUnauthorized, "wrong password")
)
