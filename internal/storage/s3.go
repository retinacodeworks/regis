package storage

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
)

type S3 struct {
	Client     *s3.Client
	BucketName string
}

func (s S3) Upload(path string, archive io.Reader) error {
	_, err := s.Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(s.BucketName),
		Key:    aws.String(""),
		Body:   archive,
	})
	return err
}

func (s S3) Has(path string) (bool, error) {
	_, err := s.Client.HeadObject(context.TODO(), &s3.HeadObjectInput{
		Bucket: aws.String(s.BucketName),
		Key:    aws.String(""),
	})
	return err != nil, err
}

func (s S3) Get(path string) (any, error) {
	_, err := s.Client.HeadObject(context.TODO(), &s3.HeadObjectInput{
		Bucket: aws.String(s.BucketName),
		Key:    aws.String(""),
	})
	return err != nil, err
}
