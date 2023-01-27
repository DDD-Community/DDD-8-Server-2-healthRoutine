package usecase

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"healthRoutine/application/domain/exercise"
)

var _ exercise.CreateHistoryUseCase = (*createHistoryUseCaseImpl)(nil)

func CreateHistoryUseCase(repo exercise.Repository, sqsCli *sqs.Client) exercise.CreateHistoryUseCase {
	return &createHistoryUseCaseImpl{
		Repository: repo,
		sqsCli:     sqsCli,
	}
}

type createHistoryUseCaseImpl struct {
	exercise.Repository
	sqsCli *sqs.Client
}

func (u *createHistoryUseCaseImpl) Use(ctx context.Context, params exercise.CreateHistoryParams) error {
	_, err := u.sqsCli.SendMessage(ctx, &sqs.SendMessageInput{
		MessageBody:  aws.String(params.UserId.String()),
		QueueUrl:     aws.String("https://sqs.ap-northeast-2.amazonaws.com/692043099242/create_history_or_drink_queue"),
		DelaySeconds: 0,
	})
	if err != nil {
		return err
	}

	return u.Repository.Create(ctx, params.UserId, params.ExerciseId, params.Weight, params.Reps, params.Set)
}
