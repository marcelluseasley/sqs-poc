package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/google/uuid"
)

type Item struct {
	name        string
	partNumber  string
	description string
}

func NewItem(name, description string) Item {
	return Item{
		name:        name,
		description: description,
		partNumber:  uuid.New().String(),
	}
}

func main() {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sqs.New(sess)

	urlResult, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String("dev-queue-1"),
	})

	queueURL := urlResult.QueueUrl

	msgResult, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            queueURL,
		MaxNumberOfMessages: aws.Int64(1),
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(msgResult.Messages) == 0 {
		fmt.Println("no more messages")
		return
	}
	fmt.Println(*msgResult.Messages[0].MessageAttributes["name"].StringValue)

}
