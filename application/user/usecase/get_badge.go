package usecase

import (
	"context"
	"github.com/google/uuid"
	"healthRoutine/application/domain/user"
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

// Use
// TODO: refactor
func (u *getBadgeUseCaseImpl) Use(ctx context.Context, userId uuid.UUID) (*user.GetBadgeResult, error) {
	resp, err := u.userRepo.GetBadgeByUserId(ctx)
	if err != nil {
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

	res, err := u.userRepo.GetLatestBadgeByUserId(ctx, userId)

	return &user.GetBadgeResult{
		MyBadge:      myBadge,
		WaitingBadge: waitingBadge,
		LatestBadge:  res.Sub,
	}, err
}
