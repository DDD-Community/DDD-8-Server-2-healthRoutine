package cmd

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

func GetAWSConfig() aws.Config {
	cfg := LoadConfig()
	return aws.Config{
		Region: cfg.AWSRegion,
		Credentials: credentials.NewStaticCredentialsProvider(
			cfg.AWSAccessKeyId,     // key id
			cfg.AWSSecretAccessKey, // secret
			"",
		),
	}
}
