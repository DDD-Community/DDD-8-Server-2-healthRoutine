package usecase

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"go.uber.org/zap"
	"healthRoutine/application/domain/user"
	"healthRoutine/pkgs/log"
	"healthRoutine/pkgs/util/format"
	"strings"
)

var _ user.UpdateProfileUseCase = (*updateProfileUseCaseImpl)(nil)

func UpdateProfileUseCase(repo user.Repository, s3Cli *s3.Client) user.UpdateProfileUseCase {
	return &updateProfileUseCaseImpl{
		Repository: repo,
		s3Cli:      s3Cli,
	}
}

type updateProfileUseCaseImpl struct {
	user.Repository
	s3Cli *s3.Client
}

func (u *updateProfileUseCaseImpl) log() *zap.SugaredLogger {
	return log.Get().Named("UPDATE_PROFILE_IMG_USE_CASE")
}

func (u *updateProfileUseCaseImpl) Use(ctx context.Context, params user.UpdateProfileParams) (err error) {
	logger := u.log()
	defer logger.Sync()

	model, _ := u.Repository.GetById(ctx, params.Id)
	if model.Nickname != params.Nickname {
		logger.Info("start check exists nickname")
		exists, ferr := u.Repository.CheckExistsByNickname(ctx, params.Nickname)
		if ferr != nil {
			logger.Error(ferr)
			return
		}

		if exists {
			err = user.ErrNicknameAlreadyExists
			logger.Error("nickname already exists")
			return
		}
	}

	logger.Info("start sorting latest file")
	lastItem, err := sortByObjectLastModified(ctx, u.s3Cli, params.Id)
	logger.Info(lastItem)
	if err != nil {
		logger.Error(err)
		return
	}

	filename := strings.Split(lastItem, "/")

	key := fmt.Sprintf("%s/%s", format.ConvertUUIDToKey(params.Id), filename[1])
	copySrc := fmt.Sprintf("%s/%s", profileTempBucketName, key)

	_, err = u.s3Cli.CopyObject(ctx, &s3.CopyObjectInput{
		Bucket:     aws.String(profileCdnBucketName),
		CopySource: &copySrc,
		Key:        &key,
	})
	if err != nil {
		logger.Error("failed to copy object")
		logger.Error(err)
		return
	}

	url := fmt.Sprintf("https://cdn.rest-api.xyz/%s", key)

	err = u.Repository.UpdateProfileById(ctx, params.Id, params.Nickname, url)
	if err != nil {
		logger.Error(err)
		return
	}
	return
}
