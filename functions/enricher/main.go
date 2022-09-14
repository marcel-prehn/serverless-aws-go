package main

import (
	"aclue.de/codetalks-lambda-enricher/service"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(event events.SQSEvent) error {
	enricherService := service.NewEnricherService()
	sqsService := service.NewSqsService()

	for _, record := range event.Records {
		models, err := enricherService.Enrich(record)
		if err != nil {
			return err
		}

		err = sqsService.Publish(models)
		if err != nil {
			return err
		}
	}
	return nil
}
