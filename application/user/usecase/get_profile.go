package usecase

import (
	"context"
	"github.com/google/uuid"
	"healthRoutine/application/domain/user"
)

var _ user.GetProfileUseCase = (*getProfileUseCaseImpl)(nil)

func GetProfileUseCase(repo user.Repository) user.GetProfileUseCase {
	return &getProfileUseCaseImpl{
		Repository: repo,
	}
}

type getProfileUseCaseImpl struct {
	user.Repository
}

func (u *getProfileUseCaseImpl) Use(ctx context.Context, id uuid.UUID) (*user.DomainModel, error) {
	return u.GetById(ctx, id)
}
