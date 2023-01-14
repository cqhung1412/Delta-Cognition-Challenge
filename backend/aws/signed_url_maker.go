package aws_util

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type AwsMaker struct {
	Config *aws.Config
	Bucket string
}

func (maker *AwsMaker) CreateS3GetSignedUrl(object_name string) (string, error) {
	client := s3.NewFromConfig(*maker.Config)
	presignClient := s3.NewPresignClient(client)

	expiration := time.Now().Add(time.Hour * 12)
	res, err := presignClient.PresignGetObject(context.Background(), &s3.GetObjectInput{
		Bucket:          aws.String("dog-recommend-012023"),
		Key:             aws.String(object_name),
		ResponseExpires: &expiration,
	})
	if err != nil {
		log.Fatal("Failed to sign request", err)
	}
	return res.URL, err
}

func (maker *AwsMaker) CreateS3UploadSignedUrl(object_name string) (string, error) {
	client := s3.NewFromConfig(*maker.Config)
	presignClient := s3.NewPresignClient(client)

	expiration := time.Hour * 2
	res, err := presignClient.PresignPutObject(context.Background(), &s3.PutObjectInput{
		Bucket: aws.String("dog-recommend-012023"),
		Key:    aws.String(object_name),
	}, s3.WithPresignExpires(expiration))
	if err != nil {
		log.Fatal("Failed to sign request", err)
	}

	time.AfterFunc(5*time.Second, func() {
		client.PutObjectAcl(context.Background(), &s3.PutObjectAclInput{
			Bucket: aws.String("dog-recommend-012023"),
			Key:    aws.String(object_name),
			ACL:    "public-read",
		})
	})

	return res.URL, err
}

func NewAwsMaker(accessKey string, secretKey string, region string, s3_bucket string) (*AwsMaker, error) {
	awsConfig, err := config.LoadDefaultConfig(
		context.Background(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
		config.WithRegion(region),
	)
	if err != nil {

		return nil, err
	}
	return &AwsMaker{
		Config: &awsConfig,
		Bucket: s3_bucket,
	}, nil
}
