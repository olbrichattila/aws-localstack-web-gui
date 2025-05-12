// Package awsshared contains AWS shared functions, like authenticate and config
package awsshared

import (
	"api/internal/database"
	"api/internal/domain"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func New(db database.Database) AWSShared {
	return &awsshared{
		db: db,
	}
}

type AWSShared interface {
	GetConfig() (*aws.Config, *context.Context, domain.Setting, error)
	GetS3Client() (*s3.Client, *context.Context, error)
	GetSQSClient() (*sqs.Client, *context.Context, error)
	GetSNSClient() (*sns.Client, *context.Context, error)
	GetDynamoDBClient() (*dynamodb.Client, *context.Context, error)
}

type awsshared struct {
	db database.Database
}

func (a *awsshared) GetConfig() (*aws.Config, *context.Context, domain.Setting, error) {
	setting, err := a.db.GetSettings()
	if err != nil {
		return nil, nil, domain.Setting{}, err
	}

	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(setting.Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			setting.Credentials.Key, setting.Credentials.Secret, "",
		)),
	)
	if err != nil {
		return nil, nil, setting, err
	}

	return &cfg, &ctx, setting, nil
}

func (s *awsshared) GetS3Client() (*s3.Client, *context.Context, error) {
	cfg, ctx, setting, err := s.GetConfig()
	if err != nil {
		return nil, nil, err
	}

	return s3.NewFromConfig(*cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(setting.Endpoint)
		o.UsePathStyle = true
	}), ctx, nil
}

// Implement GetSQSClient
func (s *awsshared) GetSQSClient() (*sqs.Client, *context.Context, error) {
	cfg, ctx, setting, err := s.GetConfig()
	if err != nil {
		return nil, nil, err
	}

	return sqs.NewFromConfig(*cfg, func(o *sqs.Options) {
		o.BaseEndpoint = aws.String(setting.Endpoint)
	}), ctx, nil
}

func (s *awsshared) GetSNSClient() (*sns.Client, *context.Context, error) {
	cfg, ctx, setting, err := s.GetConfig()
	if err != nil {
		return nil, nil, err
	}

	return sns.NewFromConfig(*cfg, func(o *sns.Options) {
		o.BaseEndpoint = aws.String(setting.Endpoint)
	}), ctx, nil
}

func (s *awsshared) GetDynamoDBClient() (*dynamodb.Client, *context.Context, error) {
	cfg, ctx, setting, err := s.GetConfig()
	if err != nil {
		return nil, nil, err
	}

	return dynamodb.NewFromConfig(*cfg, func(o *dynamodb.Options) {
		o.BaseEndpoint = aws.String(setting.Endpoint)
	}), ctx, nil
}
