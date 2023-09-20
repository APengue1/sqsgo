package session

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func New(region string) (*session.Session, error) {
	return session.NewSession(
		&aws.Config{
			Region: aws.String(region),
		},
	)
}
