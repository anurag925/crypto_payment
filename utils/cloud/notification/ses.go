package notification

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/service/ses"
)

type awsSes struct {
	*ses.SES
}

func NewAwsSes(p client.ConfigProvider, cfgs ...*aws.Config) *awsSes {
	return &awsSes{ses.New(p, cfgs...)}
}

func (a *awsSes) SendEmail(i *ses.SendEmailInput) (*ses.SendEmailOutput, error) {
	return a.SES.SendEmail(i)
}
