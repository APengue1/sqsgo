package sqsgo

import (
	"time"

	"github.com/APengue1/sqsgo/internal/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type Q struct {
	Config
	*sqs.SQS
	*sqs.ReceiveMessageInput
}

func New(config Config) (*Q, error) {
	sess, err := session.New(config.Region)
	if err != nil {
		return nil, err
	}

	q := &Q{
		Config: config,
		SQS:    sqs.New(sess),
		ReceiveMessageInput: &sqs.ReceiveMessageInput{
			QueueUrl: &config.QueueUrl,
		},
	}
	return q, nil
}

// Poll returns a channel which receives message outputs at the specified interval.
//
// Example:
//
//	config := sqsgo.Config{
//		Region:   AWS_REGION,
//		QueueUrl: AWS_SQS_QUEUE_URL,
//	}
//
// q, _ := sqsgo.New(config)
//
// interval := time.Duration(5) * time.Second
//
//	for msg := range q.Poll(interval) {
//		if msg.Error != nil {
//			log.Fatal(msg.Error)
//		}
//
//		fmt.Println(msg.ReceiveMessageOutput)
//	}
func (q *Q) Poll(interval time.Duration) <-chan ReceiveMessageOutput {
	c := make(chan ReceiveMessageOutput)
	ticker := time.NewTicker(interval)

	go func() {
		for {
			msg := q.receiveMessage()
			c <- msg
			<-ticker.C
		}
	}()

	return c
}

func (q *Q) receiveMessage() ReceiveMessageOutput {
	msgs, err := q.ReceiveMessage(q.ReceiveMessageInput)

	return ReceiveMessageOutput{
		ReceiveMessageOutput: msgs,
		Error:                err,
	}
}
