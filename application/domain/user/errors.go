package user

import "errors"

var (
	ErrEmailAlreadyExists    = errors.New("email already exists")
	ErrNicknameAlreadyExists = errors.New("nickname already exists")
)
