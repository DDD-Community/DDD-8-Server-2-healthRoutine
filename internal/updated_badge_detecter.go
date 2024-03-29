package internal

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/go-co-op/gocron"
	"github.com/google/uuid"
	"healthRoutine/application/domain/exercise"
	"healthRoutine/application/domain/user"
	"healthRoutine/application/domain/user/enum"
	"healthRoutine/pkgs/log"
	"healthRoutine/pkgs/util/dbx"
	"time"
)

const (
	updateCycle = 5
)

type SchedulerParams struct {
	UserRepo     user.Repository
	ExerciseRepo exercise.Repository
	SQSClient    *sqs.Client
}

func StartScheduler(params SchedulerParams) (err error) {
	ctx := context.Background()
	s := gocron.NewScheduler(time.UTC)
	if _, err = s.Every(updateCycle).Second().Do(task, ctx, params.UserRepo, params.ExerciseRepo, params.SQSClient); err != nil {
		return
	}

	s.StartBlocking()
	return
}

func task(
	ctx context.Context,
	userRepo user.Repository,
	exerciseRepo exercise.Repository,
	sqsCli *sqs.Client) {

	logger := log.Get()
	defer logger.Sync()

	resp, err := sqsCli.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
		QueueUrl: aws.String("https://sqs.ap-northeast-2.amazonaws.com/692043099242/create_history_or_drink_queue"),
	})
	if err != nil {
		logger.Error(err)
		return
	}

	if len(resp.Messages) == 0 {
		return
	}

	for _, v := range resp.Messages {
		userId := uuid.MustParse(*v.Body)
		logger.Info("user id: ", userId)
		var badgeId []int64
		countExercise, cerr := exerciseRepo.CountExerciseHistoryByUserId(ctx, userId)
		if cerr != nil {
			logger.Error(cerr)
		}

		countDrink, cerr := exerciseRepo.CountDrinkHistoryByUserId(ctx, userId)
		if cerr != nil {
			logger.Error(cerr)
		}
		logger.Info("total exercise: ", countExercise)
		logger.Info("total drink: ", countDrink)
		switch {
		case countExercise >= 1 && countExercise < 10:
			badgeId = append(badgeId, enum.ExerciseStart, enum.SincerityJunior)
		case countExercise >= 10 && countExercise < 30:
			badgeId = append(badgeId, enum.ExerciseHappy)
		case countExercise >= 30 && countExercise < 50:
			badgeId = append(badgeId, enum.ExerciseHolic, enum.SincerityPro)
		case countExercise >= 50 && countExercise < 100:
			badgeId = append(badgeId, enum.ExerciseMaster, enum.SincerityMaster)
		case countExercise >= 100:
			badgeId = append(badgeId, enum.ExerciseChampion, enum.SincerityChampion)
		case countDrink >= 1 && countDrink < 50:
			badgeId = append(badgeId, enum.DrinkHoneyHoney)
		case countDrink >= 50 && countDrink < 100:
			badgeId = append(badgeId, enum.DrinkBulkUpBulkUp)
		case countDrink >= 100:
			badgeId = append(badgeId, enum.DrinkHippo)
		}

		err = userRepo.CreateBadge(ctx, userId, badgeId)
		if err != nil {
			mysqlErr := dbx.UnwrapMySQLError(err)
			switch {
			case mysqlErr != nil && mysqlErr.Number == dbx.MySQLErrCodeDuplicateEntity:
				logger.Info("badge already exists for user: ", userId)
			default:
				logger.Error("failed to create badge for user: ", userId)
				logger.Error(err)
				return
			}
		}

		_, err = sqsCli.DeleteMessage(ctx, &sqs.DeleteMessageInput{
			QueueUrl:      aws.String("https://sqs.ap-northeast-2.amazonaws.com/692043099242/create_history_or_drink_queue"),
			ReceiptHandle: v.ReceiptHandle,
		})
		if err != nil {
			logger.Error("failed to delete message")
			logger.Error(err)
		}
	}

}
