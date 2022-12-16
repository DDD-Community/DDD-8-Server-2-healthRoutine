package usecase

import (
	"context"
	"healthRoutine/application/domain/user"
)

var _ user.EmailValidationUseCase = (*emailValidationUseCaseImpl)(nil)

func EmailValidationUseCase(repo user.Repository) user.EmailValidationUseCase {
	return &emailValidationUseCaseImpl{
		Repository: repo,
	}
}

type emailValidationUseCaseImpl struct {
	user.Repository
}

func (u *emailValidationUseCaseImpl) Use(ctx context.Context, email string) (err error) {
	emailExists, err := u.Repository.CheckExistsByEmail(ctx, email)
	switch {
	case emailExists:
		err = user.ErrEmailAlreadyExists
		return
	case err != nil:
		return
	}
	return
}
