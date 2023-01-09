package usecase

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"healthRoutine/application/domain/exercise"
	"healthRoutine/application/domain/user"
	"healthRoutine/pkgs/log"
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
func (u *fetchByDatetimeUseCaseImpl) Use(ctx context.Context, userId uuid.UUID, year, month int) (res exercise.FetchByDatetimeResult, err error) {
	logger := u.log()
	defer logger.Sync()

	logger.Info(year)
	logger.Info(month)

	//TODO: go routine
	resp, err := u.exerciseRepo.FetchByDateTime(ctx, userId, year, month)
	if err != nil {
		logger.Error(err)
		return
	}

	var yearOfToday int
	var monthOfToday time.Month
	var level int32
	data := make([]exercise.FetchByDatetimeDetail, 0, len(resp))
	for _, v := range resp {
		today := time.UnixMilli(v.Health.CreatedAt)
		yearOfToday = today.Year()
		monthOfToday = today.Month()

		todayCount, ferr := u.exerciseRepo.GetTodayExerciseCount(ctx, userId, v.Health.CreatedAt)
		if ferr != nil {
			logger.Error(ferr)
			return
		}

		// 다시 계산
		switch {
		case todayCount == 0:
			level = 1
		case todayCount <= 1 && todayCount >= 3:
			level = 2
		case todayCount <= 4 && todayCount >= 5:
			level = 3
		case todayCount > 5:
			level = 4
		}
		data = append(data, exercise.FetchByDatetimeDetail{
			Day:           today.Day(),
			TotalExercise: todayCount,
			Level:         level,
		})
	}

	// TODO: need sql no rows error handling
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
		welcomeMessage = fmt.Sprintf("%d%s", monthOfToday, level2)
	}

	res = exercise.FetchByDatetimeResult{
		Year:           yearOfToday,
		Month:          int(monthOfToday),
		TotalOfMonth:   len(resp),
		WelcomeMessage: welcomeMessage,
		Data:           data,
	}

	return
}
