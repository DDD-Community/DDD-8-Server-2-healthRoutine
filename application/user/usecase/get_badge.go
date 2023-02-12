package usecase

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"healthRoutine/application/domain/user"
	"healthRoutine/pkgs/log"
)

var (
	_ user.GetBadgeUseCase = (*getBadgeUseCaseImpl)(nil)
)

func GetBadgeUseCase(repo user.Repository) user.GetBadgeUseCase {
	return &getBadgeUseCaseImpl{
		userRepo: repo,
	}
}

type getBadgeUseCaseImpl struct {
	userRepo user.Repository
}

func (*getBadgeUseCaseImpl) log() *zap.SugaredLogger {
	return log.Get().Named("GET_BADGE_USE_CASE")
}

// Use
// TODO: refactor
func (u *getBadgeUseCaseImpl) Use(ctx context.Context, userId uuid.UUID) (*user.GetBadgeResult, error) {
	logger := u.log()
	defer logger.Sync()

	resp, err := u.userRepo.GetBadgeByUserId(ctx)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	// https://github.com/kyleconroy/sqlc/pull/1816
	// 고쳐지면 리팩토링
	myBadge := make([]string, 0, len(resp))
	waitingBadge := make([]string, 0, len(resp))
	for _, v := range resp {
		existsBadge, ferr := u.userRepo.ExistsBadgeByUserId(ctx, userId, v.ID)
		if ferr != nil {
			return nil, err
		}

		if !existsBadge {
			myBadge = append(myBadge, v.Sub)
		} else {
			waitingBadge = append(waitingBadge, v.Sub)
		}
	}

	var latestBadge *string
	res, ferr := u.userRepo.GetLatestBadgeByUserId(ctx, userId)
	switch {
	case ferr == user.ErrNoBadge:
		latestBadge = nil
	case ferr != nil:
		// pass
		logger.Error(err)
	default:
		latestBadge = &res.Sub
	}

	return &user.GetBadgeResult{
		MyBadge:      myBadge,
		WaitingBadge: waitingBadge,
		LatestBadge:  latestBadge,
	}, err
}
