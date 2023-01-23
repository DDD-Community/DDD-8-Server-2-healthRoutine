package usecase

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"healthRoutine/application/domain/exercise"
	"healthRoutine/application/domain/user"
	"healthRoutine/pkgs/log"
	"healthRoutine/pkgs/util/timex"
	"strconv"
	"time"
)

type welcomeByLevel string

const (
	level1 welcomeByLevel = "님 오늘도 움직여 볼까요?"
	level2 welcomeByLevel = "월은 벌컵 벌컵 하셨군요!"
)

func FetchByDatetimeUseCase(exerciseRepo exercise.Repository, userRepo user.Repository) exercise.FetchByDatetimeUseCase {
	return &fetchByDatetimeUseCaseImpl{
		exerciseRepo: exerciseRepo,
		userRepo:     userRepo,
	}
}

type fetchByDatetimeUseCaseImpl struct {
	exerciseRepo exercise.Repository
	userRepo     user.Repository
}

func (u *fetchByDatetimeUseCaseImpl) log() *zap.SugaredLogger {
	return log.Get().Named("FETCH_BY_DATETIME_USE_CASE")
}

func (u *fetchByDatetimeUseCaseImpl) Use(ctx context.Context, userId uuid.UUID, year, month int) (result exercise.FetchByDatetimeResult, err error) {
	logger := u.log()
	defer logger.Sync()

	daysInMonth := timex.GetDaysInMonth(year, time.Month(month))
	res := make([]exercise.FetchByDatetimeDetail, 0, len(daysInMonth))
	resp, err := u.exerciseRepo.FetchByDateTime(ctx, userId, year, month)
	if err != nil {
		logger.Error(err)
		return
	}

	for _, days := range daysInMonth {
		match := false
		for _, v := range resp {
			day, _ := strconv.Atoi(v.Day)
			if int(days) == day {
				res = append(res, exercise.FetchByDatetimeDetail{
					Day:           day,
					TotalExercise: v.Counts,
					Level:         getLevel(v.Counts),
					IsFutureDays:  false,
				})
				match = true
				break
			}
		}
		var level int32
		isFutureDays := false
		if time.Now().Day() < int(days) {
			isFutureDays = true
		}
		if match == false {
			if !isFutureDays {
				level = 1
			} else {
				level = 0
			}
			res = append(res, exercise.FetchByDatetimeDetail{
				Day:           int(days),
				TotalExercise: 0,
				Level:         level,
				IsFutureDays:  isFutureDays,
			})
		}
	}

	nickname, err := u.userRepo.GetNicknameById(ctx, userId)
	if err != nil {
		logger.Error(err)
		logger.Error("failed to get nickname")
		return
	}

	totalOfMonth := len(resp)
	var welcomeMessage string
	if totalOfMonth >= 30 {
		welcomeMessage = fmt.Sprintf("%s%s", nickname, level1)
	} else {
		welcomeMessage = fmt.Sprintf("%d%s", month, level2)
	}

	result = exercise.FetchByDatetimeResult{
		Year:           year,
		Month:          month,
		WelcomeMessage: welcomeMessage,
		Data:           res,
	}

	logger.Debug(result)

	return
}

func getLevel(cnt int64) (level int32) {
	switch {
	case cnt == 0:
		level = 1
	case cnt >= 1 && cnt <= 3:
		level = 2
	case cnt >= 4 && cnt <= 5:
		level = 3
	case cnt > 5:
		level = 4
	}
	return
}
