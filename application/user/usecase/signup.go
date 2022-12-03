package usecase

import (
	"context"
	"healthRoutine/application/domain/user"
)

var _ user.SignUpUseCase = (*signUpUseCaseImpl)(nil)

func SignUpUseCase(repo user.Repository) user.SignUpUseCase {
	return &signUpUseCaseImpl{
		Repository: repo,
	}
}

type signUpUseCaseImpl struct {
	user.Repository
}

func (u *signUpUseCaseImpl) Use(ctx context.Context, params user.SignUpParams) (err error) {
	return u.Repository.Create(ctx,
		params.Nickname,
		params.Email,
		params.Password)
}
