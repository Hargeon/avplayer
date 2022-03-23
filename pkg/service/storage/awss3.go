package storage

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type AWSS3 struct {
	bucketName string
	accessKey  string
	secretKey  string
	region     string
}

// NewAWSS3 initialize AWSS3
func NewAWSS3(bucketName, region, accessKey, secretKey string) *AWSS3 {
	return &AWSS3{
		bucketName: bucketName,
		accessKey:  accessKey,
		secretKey:  secretKey,
		region:     region,
	}
}

func (s *AWSS3) session() (*session.Session, error) {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(s.region),
			Credentials: credentials.NewStaticCredentials(
				s.accessKey,
				s.secretKey,
				""), // a token will be created when the session it's used.
		})

	return sess, err
}

// Download video from aws s3 to path folder, returns path to file
func (s *AWSS3) Download(ctx context.Context, id, path string) (string, error) {
	file, err := os.Create(path)

	if err != nil {
		return "", err
	}

	sess, err := s.session()

	if err != nil {
		return "", err
	}

	downloader := s3manager.NewDownloader(sess)
	_, err = downloader.DownloadWithContext(ctx, file,
		&s3.GetObjectInput{
			Bucket: aws.String(s.bucketName),
			Key:    aws.String(id),
		})

	if err != nil {
		return "", err
	}

	return path, nil
}
