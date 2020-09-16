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

	smOutput, err := svc.SendMessage(&sqs.SendMessageInput{
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"name": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String("drill set2"),
			},
			"description": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String("drill set for all your needs2"),
			},
			"partNumber": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(uuid.New().String()),
			},
		},
		MessageBody: aws.String("The message body data"),
		QueueUrl:    aws.String("https://sqs.us-east-1.amazonaws.com/327932360562/dev-queue-1"),
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(smOutput.GoString())

}
