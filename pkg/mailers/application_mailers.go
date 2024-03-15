package mailers

import (
	"bytes"
	"context"
	"fmt"
	"github.com/anurag925/crypto_payment/app"
	"github.com/anurag925/crypto_payment/utils/cloud/notification"
	"github.com/anurag925/crypto_payment/utils/cloud/session"
	"github.com/anurag925/crypto_payment/utils/logger"
	"html/template"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
)

type Client struct {
	credentials credentials
}

type credentials struct {
	host     string
	port     string
	username string
	password string
}

var (
	client     *Client
	clientOnce sync.Once
)

func instance() *Client {
	clientOnce.Do(func() {
		client = &Client{}
		client.credentials = credentials{
			host:     app.Config().SmtpHost,
			port:     app.Config().SmtpPort,
			username: app.Config().SmtpUsername,
			password: app.Config().SmtpPassword,
		}
	})
	return client
}

// Email represents the data needed for sending an email
type Email struct {
	From    string
	To      string
	Subject string
	Body    string
}

func call(to []string, from, subject, mailer_type string, data any) error {
	body, err := instance().renderTemplate(mailer_type, data)
	if err != nil {
		return err
	}
	if err := instance().sendEmail(to, from, subject, body); err != nil {
		return err
	}
	return nil
}

// renderTemplate renders the given template file with the provided data
func (c *Client) renderTemplate(templateFile string, data any) (string, error) {
	tmpl, err := template.ParseFiles(fmt.Sprintf(app.Config().DIR+"/pkg/mailers/templates/%s.html", templateFile))
	if err != nil {
		return "", err
	}

	var renderedTemplate string
	buffer := &bytes.Buffer{}
	err = tmpl.Execute(buffer, data)
	if err != nil {
		return "", err
	}

	renderedTemplate = buffer.String()
	return renderedTemplate, nil
}

// // sendEmail sends an email using the provided SMTP credentials and email data
// func (c *Client) sendEmail(to []string, from, subject, body string) error {
// 	auth := smtp.PlainAuth("", c.credentials.username, c.credentials.password, c.credentials.host)
// 	// Compose the email message
// 	msg := []byte("Subject: " + subject + "\r\n" + "\r\n" + body + "\r\n")
// 	// Send the email
// 	err := smtp.SendMail(c.credentials.host+":"+c.credentials.port, auth, from, to, msg)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// sendEmail sends an email using the provided SMTP credentials and email data
func (c *Client) sendEmail(to []string, from, subject, body string) error {
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: aws.StringSlice(to),
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Data: aws.String(body),
				},
			},
			Subject: &ses.Content{
				Data: aws.String(subject),
			},
		},
		Source: aws.String(from),
	}
	awsSession, err := session.NewAwsSession(
		app.Config().AccessKeyID, app.Config().SecretAccessKey, app.Config().Region)
	if err != nil {
		return err
	}
	res, err := notification.NewAwsSes(awsSession).SendEmail(input)
	if err != nil {
		return err
	}
	logger.Info(context.Background(), "the response is ", "res", res)
	return nil
}
