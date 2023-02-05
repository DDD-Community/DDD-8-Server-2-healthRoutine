package usecase

import (
	"context"
	"github.com/google/uuid"
	"healthRoutine/application/domain/user"
)

var _ user.WithdrawalUseCase = (*userWithdrawalUseCaseImpl)(nil)

func WithdrawalUseCase(repo user.Repository) user.WithdrawalUseCase {
	return &userWithdrawalUseCaseImpl{
		Repository: repo,
	}
}

type userWithdrawalUseCaseImpl struct {
	user.Repository
}

func (u *userWithdrawalUseCaseImpl) Use(ctx context.Context, userId uuid.UUID) error {
	return u.Repository.DeleteUserById(ctx, userId)
}
