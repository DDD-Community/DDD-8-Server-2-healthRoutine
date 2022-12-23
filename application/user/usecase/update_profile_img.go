package usecase

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"go.uber.org/zap"
	"healthRoutine/application/domain/user"
	"healthRoutine/pkgs/log"
	"healthRoutine/pkgs/util/format"
)

const (
	profileBucketName = "health-profile"
)

var _ user.UpdateProfileImgUseCase = (*updateProfileImgUseCaseImpl)(nil)

func UpdateProfileImgUseCase(repo user.Repository, s3Cli *s3.Client) user.UpdateProfileImgUseCase {
	return &updateProfileImgUseCaseImpl{
		Repository: repo,
		s3Cli:      s3Cli,
	}
}

type updateProfileImgUseCaseImpl struct {
	user.Repository
	s3Cli *s3.Client
}

func (u *updateProfileImgUseCaseImpl) log() *zap.SugaredLogger {
	return log.Get().Named("UPDATE_PROFILE_IMG_USE_CASE")
}
func (u *updateProfileImgUseCaseImpl) Use(ctx context.Context, params user.UpdateProfileImgParams) (err error) {
	logger := u.log()
	defer logger.Sync()

	key := fmt.Sprintf("%s/%s", format.ConvertUUIDToKey(params.Id), params.Filename)
	_, err = u.s3Cli.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(profileBucketName),
		Key:    &key,
		ACL:    types.ObjectCannedACLPublicRead,
		Body:   params.ProfileImg,
	})
	if err != nil {
		logger.Error(err)
		return
	}

	url := fmt.Sprintf("https://%s.s3.ap-northeast-2.amazonaws.com/%s", profileBucketName, key)

	err = u.Repository.UpdateProfileImgById(ctx, params.Id, url)
	if err != nil {
		logger.Error(err)
		return
	}
	return
}
