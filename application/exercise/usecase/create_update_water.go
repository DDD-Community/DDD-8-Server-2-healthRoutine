package usecase

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/google/uuid"
	"healthRoutine/application/domain/exercise"
)

func CreateOrUpdateWaterUseCase(repo exercise.Repository, sqsCli *sqs.Client) exercise.CreateOrUpdateWaterUseCase {
	return &createOrUpdateWaterUseCaseImpl{
		Repository: repo,
		sqsCli:     sqsCli,
	}
}

type createOrUpdateWaterUseCaseImpl struct {
	exercise.Repository
	sqsCli *sqs.Client
}

func (u *createOrUpdateWaterUseCaseImpl) Use(ctx context.Context, userId uuid.UUID, capacity int64) error {
	_, err := u.sqsCli.SendMessage(ctx, &sqs.SendMessageInput{
		MessageBody:  aws.String(userId.String()),
		QueueUrl:     aws.String("https://sqs.ap-northeast-2.amazonaws.com/692043099242/create_history_or_drink_queue"),
		DelaySeconds: 0,
	})
	if err != nil {
		return err
	}

	return u.Repository.CreateOrUpdateWater(ctx, userId, capacity)
}
