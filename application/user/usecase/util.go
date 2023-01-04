package usecase

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"healthRoutine/pkgs/util/format"
	"sort"
	"time"
)

func sortByObjectLastModified(ctx context.Context, s3Cli *s3.Client, userId uuid.UUID) (string, error) {
	resp, err := s3Cli.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(profileTempBucketName),
		Prefix: aws.String(format.ConvertUUIDToKey(userId)),
	})
	if err != nil {
		return "", err
	}

	res := make(map[string]time.Time)
	for _, v := range resp.Contents {
		res[*v.Key] = *v.LastModified
	}

	keys := make([]string, 0, len(res))
	for key := range res {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return res[keys[i]].Before(res[keys[j]])
	})

	return keys[len(keys)-1], err
}
