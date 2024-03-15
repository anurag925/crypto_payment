package session

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

type awsSession struct {
	*session.Session
}

func NewAwsSession(accessKeyId, secretAccessKey, region string) (*awsSession, error) {
	session, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKeyId, secretAccessKey, ""),
	})
	return &awsSession{session}, err
}
