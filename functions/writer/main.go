package main

import (
	"aclue.de/codetalks-lambda-writer/service"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(event events.SQSEvent) error {
	dynamodbService := service.NewDynamodbService()

	for _, record := range event.Records {
		err := dynamodbService.Save(record.Body)
		if err != nil {
			return err
		}
	}
	return nil
}
