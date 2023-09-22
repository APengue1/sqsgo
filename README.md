# SQSGO - AWS-SQS Polling, The Easy Way.

The base case for using AWS-SQS, is to send/receive messages from it.
Apart from the boilerplate required, sending is trivial.

The receiving aspect has added complexity however.
While the official aws-sdk-go is very extensive and extendable, it is very basic when it comes to receiving messages.
I do not disagree with that approach, it leaves maximum flexbility to the developer to interact with AWS via the SDK.

## The Problem

However what if you simply want to perform the base case of consuming messages from an SQS queue with little effort?
- aws-sdk-go/service/sqs is simple in that it only returns 1 or more messages, in a non-continous fashion.
- However, you are left to figure out an abstraction as to how to continuously receive messages.
- Should you use a polling strategy?
- For loop? Select? Channels?
- How about error handling?
- Should you use the main thread?

``` go
// Using the standard aws-sdk-go/service/sqs

sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(AWS_REGION),
		},
	)
if err != nil {
    log.Fatal(err)
}

q := sqs.New(sess)

// Great, we have received one message, now what?
// How do we receive the next one?
msg, err := q.ReceiveMessage(&sqs.ReceiveMessageInput{
    QueueUrl: aws.String(SQS_QUEUE_URL),
})
if err != nil {
    log.Fatal(err)
}

log.Println(msg)
```

## The Solution
``` go
config := sqsgo.Config{
		Region:       AWS_REGION,
		QueueUrl:     AWS_SQS_QUEUE_URL,
	}
q, err := sqsgo.New(config)

if err != nil {
    log.Fatal(err)
}

// Cleanly and safely receives messages continously from SQS and checks for errors.
// The polling occurs on another thread, and sends a message once it receives one.
interval := time.Duration(5) * time.Second
for msg := range q.Poll(interval) {
    if msg.Error != nil {
        log.Fatal(msg.Error)
    }

    log.Println(msg.ReceiveMessageOutput) // aws-sdk-go's *sqs.ReceiveMessageOutput
}
```
