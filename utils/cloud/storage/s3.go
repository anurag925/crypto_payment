package storage

import (
	"context"
	"fmt"
	"github.com/anurag925/crypto_payment/utils/logger"
	"io"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/service/s3"
)

type awsS3 struct {
	*s3.S3
}

var _ Client = (*awsS3)(nil)

func NewAwsS3(p client.ConfigProvider, cfgs ...*aws.Config) *awsS3 {
	return &awsS3{s3.New(p, cfgs...)}
}

func (s *awsS3) Url(bucket, key string) string {
	return fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucket, key)
}

func (s *awsS3) Upload(ctx context.Context, body io.ReadSeeker, m Metadata) error {
	logger.Info(ctx, "uploading to S3", "metadata", m)
	_, err := s.PutObject(&s3.PutObjectInput{
		ACL:         &m.ACL,
		Body:        body,
		Bucket:      &m.Bucket,
		Key:         &m.Key,
		ContentType: &m.ContentType,
	})
	return err
}

func (s *awsS3) PresignedPutUrl(ctx context.Context, m Metadata) (string, error) {
	logger.Info(ctx, "uploading to S3", "metadata", m)
	if m.ACL == "public" {
		m.ACL = s3.BucketCannedACLPublicRead
	} else {
		m.ACL = s3.BucketCannedACLPrivate
	}
	req, _ := s.PutObjectRequest(&s3.PutObjectInput{
		Bucket: &m.Bucket,
		Key:    &m.Key,
		// ACL:    &m.ACL,
	})
	return req.Presign(15 * time.Minute)
}
