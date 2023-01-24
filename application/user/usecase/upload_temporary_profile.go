package usecase

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"go.uber.org/zap"
	"golang.org/x/text/unicode/norm"
	"healthRoutine/application/domain/user"
	"healthRoutine/pkgs/log"
	"healthRoutine/pkgs/util/format"
	"strings"
	"time"
	"unicode"
)

var _ user.UploadTemporaryProfileUseCase = (*uploadTemporaryProfileUseCaseImpl)(nil)

func UploadTemporaryProfileUseCase(repo user.Repository, s3Cli *s3.Client) user.UploadTemporaryProfileUseCase {
	return &uploadTemporaryProfileUseCaseImpl{
		Repository: repo,
		s3Cli:      s3Cli,
	}
}

type uploadTemporaryProfileUseCaseImpl struct {
	user.Repository
	s3Cli *s3.Client
}

func (u *uploadTemporaryProfileUseCaseImpl) log() *zap.SugaredLogger {
	return log.Get().Named("UPLOAD_TEMP_PROFILE_USE_CASE")
}

func (u *uploadTemporaryProfileUseCaseImpl) Use(ctx context.Context, params user.UploadTemporaryProfileParams) (url string, err error) {
	logger := u.log()
	defer logger.Sync()

	var filename string
	if containsKorean(params.Filename) {
		t := time.Now().UnixMilli()
		extension := strings.Split(params.ContentType, "/")
		filename = fmt.Sprintf("%d_%d.%s", params.ContentLength, t, extension[1])
	} else {
		filename = strings.ReplaceAll(params.Filename, " ", "_")
	}

	key := fmt.Sprintf("%s/%s", format.ConvertUUIDToKey(params.Id), filename)

	_, err = u.s3Cli.PutObject(ctx, &s3.PutObjectInput{
		Bucket:        aws.String(profileTempBucketName),
		Key:           &key,
		ACL:           types.ObjectCannedACLPublicRead,
		Body:          params.ProfileImg,
		ContentType:   aws.String(params.ContentType),
		ContentLength: params.ContentLength,
	})

	if err != nil {
		logger.Error(err)
		return
	} else {

		url = fmt.Sprintf("https://%s.s3.ap-northeast-2.amazonaws.com/%s", profileTempBucketName, key)
		return
	}
}

func containsKorean(str string) bool {
	normalizedText := norm.NFC.String(str)
	for _, r := range normalizedText {
		if unicode.Is(unicode.Hangul, r) {
			return true
		}
	}
	return false
}
