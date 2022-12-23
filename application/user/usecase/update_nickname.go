package usecase

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"healthRoutine/application/domain/user"
	"healthRoutine/pkgs/log"
)

var _ user.UpdateNicknameUseCase = (*updateNicknameUseCaseImpl)(nil)

func UpdateNicknameUseCase(repo user.Repository) user.UpdateNicknameUseCase {
	return &updateNicknameUseCaseImpl{
		Repository: repo,
	}
}

type updateNicknameUseCaseImpl struct {
	user.Repository
}

func (u *updateNicknameUseCaseImpl) log() *zap.SugaredLogger {
	return log.Get().Named("UPDATE_NICKNAME_USE_CASE")
}
func (u *updateNicknameUseCaseImpl) Use(ctx context.Context, id uuid.UUID, nickname string) (err error) {
	logger := u.log()
	defer logger.Sync()

	err = u.Repository.UpdateNicknameById(ctx, id, nickname)
	if err != nil {
		logger.Error(err)
		return
	}
	return
}
