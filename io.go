package sqsgo

import "github.com/aws/aws-sdk-go/service/sqs"

type ReceiveMessageOutput struct {
	*sqs.ReceiveMessageOutput
	Error error
}
