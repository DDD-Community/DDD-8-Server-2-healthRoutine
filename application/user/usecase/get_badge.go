package usecase

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"healthRoutine/application/domain/user"
	"healthRoutine/application/domain/user/enum"
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

type getLatestBadgeUseCaseImpl struct {
	userRepo user.Repository
}

// Use
// TODO: refactor
func (u *getBadgeUseCaseImpl) Use(ctx context.Context, userId uuid.UUID) (*user.GetBadge, error) {
	resp, err := u.userRepo.GetBadgeByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	var (
		exerciseStart     bool
		exerciseHappy     bool
		exerciseHolic     bool
		exerciseMaster    bool
		exerciseChampion  bool
		sincerityJunior   bool
		sincerityPro      bool
		sincerityMaster   bool
		sincerityChampion bool
		drinkHoneyHoney   bool
		drinkBulkUpBulkUp bool
		drinkHippo        bool
	)

	for _, v := range resp {
		switch v {
		case enum.ExerciseStart:
			exerciseStart = true
		case enum.ExerciseHappy:
			exerciseHappy = true
		case enum.ExerciseHolic:
			exerciseHolic = true
		case enum.ExerciseMaster:
			exerciseMaster = true
		case enum.ExerciseChampion:
			exerciseChampion = true
		case enum.SincerityJunior:
			sincerityJunior = true
		case enum.SincerityPro:
			sincerityPro = true
		case enum.SincerityMaster:
			sincerityMaster = true
		case enum.SincerityChampion:
			sincerityChampion = true
		case enum.DrinkHoneyHoney:
			drinkHoneyHoney = true
		case enum.DrinkBulkUpBulkUp:
			drinkBulkUpBulkUp = true
		case enum.DrinkHippo:
			drinkHippo = true
		}
	}

	// TODO: go routine
	var latestBadge *user.LatestBadge
	res, err := u.userRepo.GetLatestBadgeByUserId(ctx, userId)
	if err == nil {
		latestBadge = &user.LatestBadge{
			Index:   res.ID,
			Subject: res.Subject,
		}
	} else if err == sql.ErrNoRows {
		// fix hard code
		latestBadge = nil
		err = nil
	} else {
		return nil, err
	}

	return &user.GetBadge{
		ExerciseStart:     exerciseStart,
		ExerciseHappy:     exerciseHappy,
		ExerciseHolic:     exerciseHolic,
		ExerciseMaster:    exerciseMaster,
		ExerciseChampion:  exerciseChampion,
		SincerityJunior:   sincerityJunior,
		SincerityPro:      sincerityPro,
		SincerityMaster:   sincerityMaster,
		SincerityChampion: sincerityChampion,
		DrinkHoneyHoney:   drinkHoneyHoney,
		DrinkBulkUpBulkUp: drinkBulkUpBulkUp,
		DrinkHippo:        drinkHippo,
		LatestBadge:       latestBadge,
	}, err
}
